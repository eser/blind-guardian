package updater

type UpdateConfig struct {
	OutputDir    string `conf:"OUTPUT_DIR"    default:"etc/tempmail_db"`
	OutputFile   string `conf:"OUTPUT_FILE"   default:"%s_domains.txt"`
	BlocklistUri string `conf:"BLOCKLIST_URI" default:"https://raw.githubusercontent.com/disposable/disposable-email-domains/refs/heads/master/domains.txt"` //nolint:lll
}
