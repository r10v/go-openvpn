package primitives

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	log "github.com/cihub/seelog"
	"os"
)

// CreateTLSCryptKey generates symmetric key in HEX format 2048 bits length
func (sp *SecurityPrimitives) CreateTLSCryptKey() error {
	taKey := make([]byte, 256)
	_, err := rand.Read(taKey)
	if err != nil {
		fmt.Println("error:", err)
		return err
	}

	keyOut, err := os.OpenFile(sp.TLSCryptKeyPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Info("failed to open "+sp.TLSCryptKeyPath+" for writing:", err)
		return err
	}
	defer keyOut.Close()

	var keyEntries []string
	keyEntries = append(keyEntries, "-----BEGIN OpenVPN Static key V1-----\n")
	keyEntries = append(keyEntries, hex.EncodeToString(taKey))
	keyEntries = append(keyEntries, "\n-----END OpenVPN Static key V1-----\n")

	for _, s := range keyEntries {
		_, err := keyOut.WriteString(s)
		if err != nil {
			return err
		}
	}

	log.Debug("written " + sp.TLSCryptKeyPath)

	return nil
}
