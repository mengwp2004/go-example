

/ crypto/rand.Reader is a good source of entropy for blinding the RSA
// operation.
rng := rand.Reader

message := []byte("message to be signed")

// Only small messages can be signed directly; thus the hash of a
// message, rather than the message itself, is signed. This requires
// that the hash function be collision resistant. SHA-256 is the
// least-strong hash function that should be used for this at the time
// of writing (2016).
hashed := sha256.Sum256(message)

signature, err := SignPKCS1v15(rng, rsaPrivateKey, crypto.SHA256, hashed[:])
if err != nil {
        fmt.Fprintf(os.Stderr, "Error from signing: %s\n", err)
        return
}

fmt.Printf("Signature: %x\n", signature)
