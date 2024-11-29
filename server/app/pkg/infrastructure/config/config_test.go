package config

import (
	"os"
	"testing"

	"gopkg.in/yaml.v3"
)

var okConfig = &Config{
	Stage: "dev",
	Db: ConfigDb{
		Host:     "localhost",
		Port:     3306,
		User:     "root",
		Password: "password",
		Database: "test",
	},
}

func TestLoadConfig(t *testing.T) {
	// テスト用の設定ファイルを作成
	configContent, err := yaml.Marshal(okConfig)
	if err != nil {
		t.Fatalf("設定ファイルの作成に失敗しました: %v", err)
	}

	file, err := os.CreateTemp("", "test_config.yaml")
	if err != nil {
		t.Fatalf("一時ファイルの作成に失敗しました: %v", err)
	}
	defer os.Remove(file.Name())

	if _, err := file.Write(configContent); err != nil {
		t.Fatalf("一時ファイルへの書き込みに失敗しました: %v", err)
	}
	if err := file.Close(); err != nil {
		t.Fatalf("一時ファイルのクローズに失敗しました: %v", err)
	}

	// LoadConfig 関数をテスト
	config, err := LoadConfig(file.Name())
	if err != nil {
		t.Fatalf("LoadConfig 関数の実行に失敗しました: %v", err)
	}

	if config.Stage != okConfig.Stage {
		t.Errorf("期待される Stage: %s, 実際の Stage: %s", okConfig.Stage, config.Stage)
	}
}

var ngConfigList = []Config{
	Config{},
	Config{
		Stage: "",
		Db:    ConfigDb{},
	},
}

func TestLoadConfigValidationFailure(t *testing.T) {
	// バリデーションに失敗する設定ファイルを作成

	for _, badConfig := range ngConfigList {
		invalidConfigContent, err := yaml.Marshal(badConfig)
		if err != nil {
			t.Fatalf("設定ファイルの作成に失敗しました: %v", err)
		}

		file, err := os.CreateTemp("", "invalid_test_config.yaml")
		if err != nil {
			t.Fatalf("一時ファイルの作成に失敗しました: %v", err)
		}
		defer os.Remove(file.Name())

		if _, err := file.Write([]byte(invalidConfigContent)); err != nil {
			t.Fatalf("一時ファイルへの書き込みに失敗しました: %v", err)
		}
		if err := file.Close(); err != nil {
			t.Fatalf("一時ファイルのクローズに失敗しました: %v", err)
		}

		// LoadConfig 関数をテスト
		_, err = LoadConfig(file.Name())
		if err == nil {
			t.Fatalf("バリデーションに失敗するはずの設定ファイルでエラーが発生しませんでした")
		}

	}
}
