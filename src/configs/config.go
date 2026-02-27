package configs

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

// Config struktur konfigurasi utama
type Config struct {
	Server struct {
		IP    string `yaml:"ip" json:"ip"`
		Port  int    `yaml:"port" json:"port"`
		Token string `json:"token"`
		Auth  struct {
			Store struct {
				Type   string `yaml:"type" json:"type"`     // memory/file/redis
				Expiry int    `yaml:"expiry" json:"expiry"` // Waktu kedaluwarsa (jam)
			} `yaml:"store" json:"store"`
		} `yaml:"auth" json:"auth"`
		ServerVersion string `yaml:"server_version" json:"server_version"`
	} `yaml:"server" json:"server"`

	// Konfigurasi lapisan transport
	Transport struct {
		WebSocket struct {
			Enabled bool   `yaml:"enabled" json:"enabled"`
			IP      string `yaml:"ip" json:"ip"`
			Port    int    `yaml:"port" json:"port"`
		} `yaml:"websocket" json:"websocket"`

		MQTTUDP struct {
			Enabled bool `yaml:"enabled" json:"enabled"`
			MQTT    struct {
				IP   string `yaml:"ip" json:"ip"`
				Port int    `yaml:"port" json:"port"`
				QoS  int    `yaml:"qos" json:"qos"`
			} `yaml:"mqtt" json:"mqtt"`
			UDP struct {
				IP                string `yaml:"ip" json:"ip"`
				ShowPort          int    `yaml:"show_port" json:"show_port"` // Port tampilan
				Port              int    `yaml:"port" json:"port"`
				SessionTimeout    string `yaml:"session_timeout" json:"session_timeout"`
				MaxPacketSize     int    `yaml:"max_packet_size" json:"max_packet_size"`
				EnableReliability bool   `yaml:"enable_reliability" json:"enable_reliability"`
			} `yaml:"udp" json:"udp"`
		} `yaml:"mqtt_udp" json:"mqtt_udp"`
	} `yaml:"transport" json:"transport"`

	Log struct {
		LogLevel string `yaml:"log_level" json:"log_level"`
		LogDir   string `yaml:"log_dir" json:"log_dir"`
		LogFile  string `yaml:"log_file" json:"log_file"`
	} `yaml:"log" json:"log"`

	Web struct {
		Port         int    `yaml:"port" json:"port"`
		StaticDir    string `yaml:"static_dir" json:"static_dir"`
		Websocket    string `yaml:"websocket" json:"websocket"`
		VisionURL    string `yaml:"vision" json:"vision"`
		ActivateText string `yaml:"activate_text" json:"activate_text"` // Teks yang disertakan saat mengirim kode aktivasi
	} `yaml:"web" json:"web"`

	DefaultPrompt   string        `yaml:"prompt"             json:"prompt"`
	Roles           []Role        `yaml:"roles"              json:"roles"` // Daftar peran
	DeleteAudio     bool          `yaml:"delete_audio"       json:"delete_audio"`
	QuickReply      bool          `yaml:"quick_reply"        json:"quick_reply"`
	QuickReplyWords []string      `yaml:"quick_reply_words"  json:"quick_reply_words"`
	LocalMCPFun     []LocalMCPFun `yaml:"local_mcp_fun"      json:"local_mcp_fun"` // Pemetaan fungsi MCP lokal
	SaveTTSAudio    bool          `yaml:"save_tts_audio"  json:"save_tts_audio"`   // Apakah file audio TTS disimpan
	SaveUserAudio   bool          `yaml:"save_user_audio" json:"save_user_audio"`  // Apakah file audio pengguna disimpan

	SelectedModule map[string]string `yaml:"selected_module" json:"selected_module"`

	PoolConfig    PoolConfig    `yaml:"pool_config"`
	McpPoolConfig McpPoolConfig `yaml:"mcp_pool_config"`

	ASR   map[string]ASRConfig  `yaml:"ASR"   json:"ASR"`
	TTS   map[string]TTSConfig  `yaml:"TTS"   json:"TTS"`
	LLM   map[string]LLMConfig  `yaml:"LLM"   json:"LLM"`
	VLLLM map[string]VLLMConfig `yaml:"VLLLM" json:"VLLLM"`

	CMDExit []string `yaml:"CMD_exit" json:"CMD_exit"`
}

type LocalMCPFun struct {
	Name        string `yaml:"name"         json:"name"`        // Nama fungsi
	Description string `yaml:"description"  json:"description"` // Deskripsi fungsi
	Enabled     bool   `yaml:"enabled"      json:"enabled"`     // Apakah diaktifkan
}

type Role struct {
	Name        string `yaml:"name"         json:"name"`        // Nama peran
	Description string `yaml:"description"  json:"description"` // Deskripsi peran
	Enabled     bool   `yaml:"enabled"      json:"enabled"`     // Apakah diaktifkan
}

type PoolConfig struct {
	PoolMinSize       int `yaml:"pool_min_size"`
	PoolMaxSize       int `yaml:"pool_max_size"`
	PoolRefillSize    int `yaml:"pool_refill_size"`
	PoolCheckInterval int `yaml:"pool_check_interval"`
}
type McpPoolConfig struct {
	PoolMinSize       int `yaml:"pool_min_size"`
	PoolMaxSize       int `yaml:"pool_max_size"`
	PoolRefillSize    int `yaml:"pool_refill_size"`
	PoolCheckInterval int `yaml:"pool_check_interval"`
}

// ASRConfig struktur konfigurasi ASR
type ASRConfig map[string]interface{}

type VoiceInfo struct {
	Name        string `yaml:"name"         json:"name"`         // Nama suara, sesuai dengan string timbre TTS, mis. zh_female_wanwanxiaohe_moon_bigtts
	Language    string `yaml:"language"     json:"language"`     // Bahasa, menandai jenis bahasa, digunakan untuk pemilihan frontend
	DisplayName string `yaml:"display_name" json:"display_name"` // Nama tampilan untuk frontend, mis. WanWan XiaoHe
	Sex         string `yaml:"sex"          json:"sex"`          // Jenis kelamin, Laki-laki/Perempuan
	Description string `yaml:"description"  json:"description"`  // Deskripsi timbre suara
	AudioURL    string `yaml:"audio_url"    json:"audio_url"`    // URL audio untuk pratinjau
}

// TTSConfig struktur konfigurasi TTS
type TTSConfig struct {
	Type            string      `yaml:"type"             json:"type"`             // Tipe TTS
	Voice           string      `yaml:"voice"            json:"voice"`            // Nama suara
	Format          string      `yaml:"format"           json:"format"`           // Format output
	OutputDir       string      `yaml:"output_dir"       json:"output_dir"`       // Direktori output
	AppID           string      `yaml:"appid"            json:"appid"`            // ID aplikasi
	Token           string      `yaml:"token"            json:"token"`            // Kunci API
	Cluster         string      `yaml:"cluster"          json:"cluster"`          // Informasi cluster
	SupportedVoices []VoiceInfo `yaml:"supported_voices" json:"supported_voices"` // Daftar suara yang didukung
}

// LLMConfig struktur konfigurasi LLM
type LLMConfig struct {
	Type        string                 `yaml:"type"        json:"type"`        // Tipe LLM
	ModelName   string                 `yaml:"model_name"  json:"model_name"`  // Nama model
	BaseURL     string                 `yaml:"url"         json:"url"`         // Alamat API
	APIKey      string                 `yaml:"api_key"     json:"api_key"`     // Kunci API
	Temperature float64                `yaml:"temperature" json:"temperature"` // Parameter suhu
	MaxTokens   int                    `yaml:"max_tokens"  json:"max_tokens"`  // Jumlah token maksimum
	TopP        float64                `yaml:"top_p"       json:"top_p"`       // Parameter TopP
	Extra       map[string]interface{} `yaml:",inline"     json:"extra"`       // Konfigurasi tambahan
}

// SecurityConfig struktur konfigurasi keamanan gambar
type SecurityConfig struct {
	MaxFileSize       int64    `yaml:"max_file_size"      json:"max_file_size"`      // Ukuran file maksimum (byte)
	MaxPixels         int64    `yaml:"max_pixels"         json:"max_pixels"`         // Jumlah piksel maksimum
	MaxWidth          int      `yaml:"max_width"          json:"max_width"`          // Lebar maksimum
	MaxHeight         int      `yaml:"max_height"         json:"max_height"`         // Tinggi maksimum
	AllowedFormats    []string `yaml:"allowed_formats"    json:"allowed_formats"`    // Format gambar yang diizinkan
	EnableDeepScan    bool     `yaml:"enable_deep_scan"   json:"enable_deep_scan"`   // Aktifkan pemindaian keamanan mendalam
	ValidationTimeout string   `yaml:"validation_timeout" json:"validation_timeout"` // Waktu tunggu validasi
}

// VLLMConfig struktur konfigurasi VLLLM (Model Bahasa Visual Besar)
type VLLMConfig struct {
	Type        string                 `yaml:"type"        json:"type"`        // Tipe API, menggunakan tipe LLM
	ModelName   string                 `yaml:"model_name"  json:"model_name"`  // Nama model yang mendukung visual
	BaseURL     string                 `yaml:"url"         json:"url"`         // Alamat API
	APIKey      string                 `yaml:"api_key"     json:"api_key"`     // Kunci API
	Temperature float64                `yaml:"temperature" json:"temperature"` // Parameter suhu
	MaxTokens   int                    `yaml:"max_tokens"  json:"max_tokens"`  // Jumlah token maksimum
	TopP        float64                `yaml:"top_p"       json:"top_p"`       // Parameter TopP
	Security    SecurityConfig         `yaml:"security"    json:"security"`    // Konfigurasi keamanan gambar
	Extra       map[string]interface{} `yaml:",inline"     json:"extra"`       // Konfigurasi tambahan
}

var (
	Cfg *Config
)

func (cfg *Config) ToString() string {
	data, _ := yaml.Marshal(cfg)
	return string(data)
}

func (cfg *Config) FromString(data string) error {
	return yaml.Unmarshal([]byte(data), cfg)
}

func (cfg *Config) SaveToDB(dbi ConfigDBInterface) error {
	data := cfg.ToString()
	return dbi.UpdateServerConfig(data)
}

// LoadConfig memuat konfigurasi
// Sepenuhnya memuat konfigurasi dari database; jika database kosong, gunakan konfigurasi default dan inisialisasi database
func LoadConfig(dbi ConfigDBInterface) (*Config, string, error) {
	bUseDatabaseCfg := true
	// Coba muat konfigurasi dari database
	cfgStr, err := dbi.LoadServerConfig()
	if err != nil {
		fmt.Println("Gagal memuat konfigurasi server:", err)
		return nil, "", err
	}

	config := &Config{}

	path := "database:serverConfig"
	if cfgStr != "" {
		config.FromString(cfgStr)
		LoadProvidersFromDB(dbi, config)
		config = CheckAndModifyConfig(config)
		Cfg = config
		if bUseDatabaseCfg {
			return Cfg, path, nil
		}
	}

	// Coba baca dari file
	path = ".config.yaml"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		path = "config.yaml"
	}

	data, err := os.ReadFile(path)
	if err != nil {
		// Gagal membaca file konfigurasi, gunakan konfigurasi default
		config = NewDefaultInitConfig()
		data, _ = yaml.Marshal(config)
	} else {
		if err := yaml.Unmarshal(data, config); err != nil {
			return nil, path, err
		}
	}

	err = dbi.InitServerConfig(string(data))
	if err != nil {
		fmt.Println("Gagal menginisialisasi konfigurasi server ke database:", err)
	}
	config = CheckAndModifyConfig(config)
	Cfg = config
	return config, path, nil
}

func LoadProvidersFromDB(dbi ConfigDBInterface, cfg *Config) {
	// Muat konfigurasi provider ASR
	asrData := dbi.LoadProviderData("ASR", 0)
	if asrData != nil {
		//fmt.Println("ASR Providers:", asrData)
		cfg.ASR = make(map[string]ASRConfig)
		for name, dataStr := range asrData {
			var asrConfig ASRConfig
			if err := yaml.Unmarshal([]byte(dataStr), &asrConfig); err == nil {
				cfg.ASR[name] = asrConfig
			}
		}
	}
	// Muat konfigurasi provider TTS
	ttsData := dbi.LoadProviderData("TTS", 0)
	if ttsData != nil {
		//fmt.Println("TTS Providers:", ttsData)
		cfg.TTS = make(map[string]TTSConfig)
		for name, dataStr := range ttsData {
			var ttsConfig TTSConfig
			if err := yaml.Unmarshal([]byte(dataStr), &ttsConfig); err == nil {
				cfg.TTS[name] = ttsConfig
			}
		}
	}
	// Muat konfigurasi provider LLM
	llmData := dbi.LoadProviderData("LLM", 0)
	if llmData != nil {
		//fmt.Println("LLM Providers:", llmData)
		cfg.LLM = make(map[string]LLMConfig)
		for name, dataStr := range llmData {
			var llmConfig LLMConfig
			if err := yaml.Unmarshal([]byte(dataStr), &llmConfig); err == nil {
				cfg.LLM[name] = llmConfig
			}
		}
	}
	// Muat konfigurasi provider VLLLM
	vllmData := dbi.LoadProviderData("VLLLM", 0)
	if vllmData != nil {
		//fmt.Println("VLLLM Providers:", vllmData)
		cfg.VLLLM = make(map[string]VLLMConfig)
		for name, dataStr := range vllmData {
			var vllmConfig VLLMConfig
			if err := yaml.Unmarshal([]byte(dataStr), &vllmConfig); err == nil {
				cfg.VLLLM[name] = vllmConfig
			}
		}
	}
	dbi.UpdateServerConfig(cfg.ToString())
}

func CheckAndModifyConfig(cfg *Config) *Config {
	// Periksa Cfg.LocalMCPFun semua huruf kecil dan hapus spasi
	if cfg.LocalMCPFun == nil {
		cfg.LocalMCPFun = []LocalMCPFun{}
	}
	fmt.Printf("Memeriksa konfigurasi: LocalMCPFun cnt %d\n", len(cfg.LocalMCPFun))
	if len(cfg.LocalMCPFun) < 10 {
		for i := 0; i < len(cfg.LocalMCPFun); i++ {
			cfg.LocalMCPFun[i].Name = strings.ToLower(strings.TrimSpace(cfg.LocalMCPFun[i].Name))
			cfg.LocalMCPFun[i].Description = strings.ToLower(strings.TrimSpace(cfg.LocalMCPFun[i].Description))
		}
	}
	// Periksa apakah konfigurasi default ASR, LLM, TTS dan VLLLM ada
	if cfg.SelectedModule == nil {
		cfg.SelectedModule = map[string]string{}
	}
	if cfg.LLM == nil {
		cfg.LLM = map[string]LLMConfig{}
	}
	if cfg.VLLLM == nil {
		cfg.VLLLM = map[string]VLLMConfig{}
	}
	if cfg.ASR == nil {
		cfg.ASR = map[string]ASRConfig{}
	}
	if cfg.TTS == nil {
		cfg.TTS = map[string]TTSConfig{}
	}
	fmt.Printf("Memeriksa konfigurasi: LLM:%d, VLLLM:%d, ASR:%d, TTS:%d\n", len(cfg.LLM), len(cfg.VLLLM), len(cfg.ASR), len(cfg.TTS))
	fmt.Println("Memeriksa konfigurasi: SelectedModule", cfg.SelectedModule)
	// Jika SelectedModule tidak dipilih atau yang dipilih tidak ada, pilih yang pertama
	llmName, ok := cfg.SelectedModule["LLM"]
	_, exists := cfg.LLM[llmName]
	if !ok || llmName == "" || !exists {
		// Pilih LLM yang tersedia sebagai default
		for name := range cfg.LLM {
			cfg.SelectedModule["LLM"] = name
			fmt.Println("LLM default tidak diatur atau LLM yang diatur tidak ada, telah diatur ke", name)
			break
		}
	}
	defaulCfg := NewDefaultInitConfig()
	if len(cfg.LLM) == 0 {
		fmt.Println("Peringatan: Tidak ada provider LLM yang tersedia, menggunakan konfigurasi default!")
		cfg.LLM = defaulCfg.LLM
		for name := range cfg.LLM {
			cfg.SelectedModule["LLM"] = name
			fmt.Println("LLM default telah diatur ke", name)
			break
		}
	}

	vlllmName, ok := cfg.SelectedModule["VLLLM"]
	_, exists = cfg.VLLLM[vlllmName]
	if !ok || vlllmName == "" || !exists {
		// Pilih VLLLM yang tersedia sebagai default
		for name := range cfg.VLLLM {
			cfg.SelectedModule["VLLLM"] = name
			fmt.Println("VLLLM default tidak diatur atau VLLLM yang diatur tidak ada, telah diatur ke", name)
			break
		}
	}
	if len(cfg.VLLLM) == 0 {
		fmt.Println("Peringatan: Tidak ada provider VLLLM yang tersedia, menggunakan konfigurasi default!")
		cfg.VLLLM = defaulCfg.VLLLM
		for name := range cfg.VLLLM {
			cfg.SelectedModule["VLLLM"] = name
			fmt.Println("VLLLM default telah diatur ke", name)
			break
		}
	}

	asrName, ok := cfg.SelectedModule["ASR"]
	_, exists = cfg.ASR[asrName]
	// ASRConfig adalah map[string]interface{}, hanya periksa apakah key ada dan name tidak kosong
	if !ok || asrName == "" || !exists {
		// Pilih ASR yang tersedia sebagai default
		for name := range cfg.ASR {
			cfg.SelectedModule["ASR"] = name
			fmt.Println("ASR default tidak diatur atau ASR yang diatur tidak ada, telah diatur ke", name)
			break
		}
	}

	if len(cfg.ASR) == 0 {
		fmt.Println("Peringatan: Tidak ada provider ASR yang tersedia, menggunakan konfigurasi default!")
		cfg.ASR = defaulCfg.ASR
		for name := range cfg.ASR {
			cfg.SelectedModule["ASR"] = name
			fmt.Println("ASR default telah diatur ke", name)
			break
		}
	}

	ttsName, ok := cfg.SelectedModule["TTS"]
	_, exists = cfg.TTS[ttsName]
	if !ok || ttsName == "" || !exists {
		// Pilih TTS yang tersedia sebagai default
		for name := range cfg.TTS {
			cfg.SelectedModule["TTS"] = name
			fmt.Println("TTS default tidak diatur atau TTS yang diatur tidak ada, telah diatur ke", name)
			break
		}
	}

	if len(cfg.TTS) == 0 {
		fmt.Println("Peringatan: Tidak ada provider TTS yang tersedia, menggunakan konfigurasi default!")
		cfg.TTS = defaulCfg.TTS
		for name := range cfg.TTS {
			cfg.SelectedModule["TTS"] = name
			fmt.Println("TTS default telah diatur ke", name)
			break
		}
	}

	return cfg
}
