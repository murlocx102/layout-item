package configs

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"

	"github.com/spf13/viper"
)

var (
	// 默认目录列表 configs conf data
	configDir = []string{"configs", "conf", "data"}
	// 默认文件名列表 custom:个人 test:测试 production:正式
	configNames = []string{"custom", "test", "production"}
)

var config *BaseConfig

// GetConfig 获取配置
func GetConfig() *BaseConfig {
	return config
}

/** LoadConfing 加载解析配置
  参数:
  *       conf            interface{}	反序列化对象,必须为非空指针
  *       dir             string     	配置文件目录(空串就依次查询默认.直到符合为止)
  *       fileName        string     	配置文件名(空串就依次查询默认.直到符合为止)
  返回值:
  *       error   error
*/
func LoadConfing(dir, fileName string) error {
	cfg := &BaseConfig{}
	if err := mustNotNilPtr(cfg); err != nil {
		return err
	}

	var (
		filePath string
		err      error
		v        *viper.Viper
	)

	switch {
	case dir != "":
		filePath, err = existFilePath(dir)
		if err != nil {
			filePath, err = getDefaultFilePath()
			if err != nil {
				break
			}
		}

		v, err = readConfFile(filePath, fileName)

	case dir == "":
		filePath, err = getDefaultFilePath()
		if err != nil {
			break
		}

		v, err = readConfFile(filePath, fileName)

	default:
		return errors.New("加载配置文件异常")
	}

	if err != nil {
		return err
	}

	if err := v.Unmarshal(cfg); err != nil {
		return err
	}

	config = cfg

	return nil
}

// readConfFile 读取配置文件
func readConfFile(dir, fileName string) (*viper.Viper, error) {
	v := viper.New()

	if fileName == "" {
		fs, err := os.ReadDir(dir)
		if err != nil {
			return nil, fmt.Errorf("配置文件目录异常:%s", dir)
		}

		// 不进行递归目录查询
		for _, f := range fs {
			if f.IsDir() {
				continue
			} else {
				for _, value := range configNames {
					name := removeFileSuffix(f.Name())
					if name == value {
						// 仅读取匹配默认的第一个
						fileName = f.Name()
						break
					}
				}
			}
		}

		if fileName == "" {
			return nil, fmt.Errorf("未找到默认配置文件的任意一个:%v,请指定配置文件", configNames)
		}
	}

	name := removeFileSuffix(fileName)

	v.SetConfigName(name)
	// 可多目录
	v.AddConfigPath(dir)

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("读取[%s]文件失败,原因:%s", fileName, err.Error())
	}

	return v, nil
}

// removeFileSuffix 去除文件名后缀
func removeFileSuffix(fileName string) string {
	name := fileName
	// 有后缀去掉
	ext := filepath.Ext(fileName)
	if ext != "" {
		fb := filepath.Base(fileName)

		name = fb[0 : len(fb)-len(ext)]
	}

	return name
}

/** existFilePath 指定路径是否存在
  参数:
  *       dir     string	非空路径 (支持相对和绝对路径)
  返回值:
  *       string  string	存在则返回绝对路径
  *       error   error
*/
func existFilePath(dir string) (string, error) {
	fp, err := absolutePath(dir)
	if err != nil {
		return "", err
	}

	info, err := os.Stat(fp)
	if err != nil && errors.Is(err, os.ErrNotExist) {
		return "", errors.New("指定路径不存在")
	}

	if !info.IsDir() {
		return "", fmt.Errorf("[%s] 非目录", dir)
	}

	return fp, nil
}

// absolutePath 构建绝对路径
func absolutePath(dir string) (string, error) {
	if !filepath.IsAbs(dir) {
		fp, err := filepath.Abs(dir)
		if err != nil {
			return "", errors.New("获取系统路径失败,构建绝对路径失败")
		}

		return fp, nil
	}

	return dir, nil
}

// getDefaultFilePath 获取默认路径 依次判定: config -> conf -> data
func getDefaultFilePath() (string, error) {
	for _, dir := range configDir {
		fp, err := existFilePath(dir)
		if err != nil {
			continue
		}

		return fp, nil
	}

	return "", fmt.Errorf("未找到符合的默认目录 %v", configDir)
}

// mustNotNilPtr 解析对象不为空,且非空指针
func mustNotNilPtr(conf interface{}) error {
	if reflect.TypeOf(conf).Kind() != reflect.Ptr {
		return errors.New("conf参数必须是指针")
	}

	if reflect.ValueOf(conf).IsNil() {
		return errors.New("conf不能为nil")
	}

	return nil
}
