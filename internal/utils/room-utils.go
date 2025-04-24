package utils

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math"
)

// GenerateRoomID creates a room ID using the hostID and 16 random bytes.
func GenerateRoomID(hostID string) (string, error) {
    totalLen := 36

    if len(hostID) >= totalLen {
        return "", fmt.Errorf("hostID too long; must be less than %d characters", totalLen)
    }

    // One character is reserved for the dash; the remainder is the salted hex string.
    needed := totalLen - len(hostID) - 1
    if needed <= 0 {
        return "", fmt.Errorf("invalid configuration: total length too short to combine hostID with separator")
    }

    // Generate enough random bytes so that when converted to hex, they provide at least 'needed' characters.
    numBytes := int(math.Ceil(float64(needed) / 2))
    randomBytes := make([]byte, numBytes)
    if _, err := rand.Read(randomBytes); err != nil {
        return "", err
    }

    // Use hostID as salt in HMAC-SHA256.
    h := hmac.New(sha256.New, []byte(hostID))
    h.Write(randomBytes)
    saltedDigest := h.Sum(nil)
    saltedHex := hex.EncodeToString(saltedDigest)

    // Truncate the hex string to exactly 'needed' characters.
    if len(saltedHex) > needed {
        saltedHex = saltedHex[:needed]
    }

    roomID := fmt.Sprintf("%s-%s", hostID, saltedHex)
    if len(roomID) != totalLen {
        return "", fmt.Errorf("generated roomID length is %d; expected %d", len(roomID), totalLen)
    }

    return roomID, nil
}
