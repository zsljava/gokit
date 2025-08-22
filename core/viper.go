package core

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/zsljava/gokit/global"
	"github.com/zsljava/gokit/util/aes"
	"os"
	"regexp"
	"strings"
)

func NewConfig(p string) *viper.Viper {
	envConf := os.Getenv("APP_CONF")
	if envConf == "" {
		envConf = p
	}
	fmt.Println("load conf file:", envConf)
	return getConfig(envConf)
}

func getConfig(path string) *viper.Viper {
	conf := viper.New()
	conf.SetConfigFile(path)
	err := conf.ReadInConfig()
	if err != nil {
		panic(err)
	}

	// 环境变量中获取解密的值
	ihrConfigSecret := os.Getenv("CONFIG_SECRET")
	// 解密所有被 Encrypt(...) 包裹的值
	decryptedConfig, err := decryptConfigValues(ihrConfigSecret, conf.AllSettings())
	if err != nil {
		panic(fmt.Errorf("decrypt config fail: %v", err))
	}

	// 将解密后的配置重新设置回viper
	for key, value := range decryptedConfig.(map[string]interface{}) {
		conf.Set(key, value)
	}

	conf.WatchConfig()

	conf.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = conf.Unmarshal(&global.CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err = conf.Unmarshal(&global.CONFIG); err != nil {
		panic(err)
	}
	global.VIPER = conf
	return conf
}

// 判断是否是加密字段
func isEncryptedField(value string) bool {
	return strings.HasPrefix(value, `Encrypt(`) && strings.HasSuffix(value, `)`)
}

// 从加密字段中提取加密内容
func extractEncryptedContent(encryptedStr string) string {
	// 使用正则提取 Encrypt("...") 中的内容
	re := regexp.MustCompile(`^Encrypt\((.*)\)$`)
	matches := re.FindStringSubmatch(encryptedStr)
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}

// 实际解密函数 (使用AES示例)
func decryptValue(ihrConfigSecret, encryptedContent string) (string, error) {
	if ihrConfigSecret == "" {
		return "", fmt.Errorf("CONFIG_SECRET environment variable not set")
	}
	// 这里实现你的解密逻辑，比如AES解密
	// 示例:
	decrypt := aes.Decrypt(encryptedContent, ihrConfigSecret)
	return decrypt, nil // 示例中直接返回，实际使用时替换为真实解密逻辑
}

func decryptConfigValues(ihrConfigSecret string, config interface{}) (interface{}, error) {
	switch v := config.(type) {
	case map[string]interface{}:
		for key, value := range v {
			decrypted, err := decryptConfigValues(ihrConfigSecret, value)
			if err != nil {
				return nil, err
			}
			v[key] = decrypted
		}
		return v, nil
	case []interface{}:
		for i, item := range v {
			decrypted, err := decryptConfigValues(ihrConfigSecret, item)
			if err != nil {
				return nil, err
			}
			v[i] = decrypted
		}
		return v, nil
	case string:
		if isEncryptedField(v) {
			encryptedContent := extractEncryptedContent(v)
			decrypted, err := decryptValue(ihrConfigSecret, encryptedContent)
			if err != nil {
				return nil, fmt.Errorf("decrypt config fail: %v", err)
			}
			return decrypted, nil
		}
		return v, nil
	default:
		return v, nil
	}
}
