# Chat App Backend

A secure, real-time chat application backend built with Go, featuring WebSocket connections and end-to-end encryption using AES and Diffie-Hellman key exchange.

## 🚀 Features

- **Real-time Communication**: WebSocket-based messaging for instant communication
- **End-to-End Encryption**: Secure message encryption using AES-256 in CTR mode
- **Diffie-Hellman Key Exchange**: Secure key establishment between clients
- **Lightweight Architecture**: Simple and efficient Go-based server
- **Cross-Origin Support**: CORS-enabled for web client integration

## 🔐 Security Features

### AES Encryption
- **Algorithm**: AES-256 in Counter (CTR) mode
- **Key Size**: 256-bit keys derived from Diffie-Hellman shared secrets
- **IV Generation**: Cryptographically secure random IV for each message

### Diffie-Hellman Key Exchange
- **Prime**: 2048-bit safe prime for secure key exchange
- **Generator**: 2 (standard DH generator)
- **Key Derivation**: SHA-256 hash of shared secret for AES key

## 🛠️ Installation

### Prerequisites
- Go 1.24.4 or higher
- Git

### Setup
1. Clone the repository:
```bash
git clone <repository-url>
cd chat-app-backend
```

2. Install dependencies:
```bash
go mod download
```

3. Run the server:
```bash
go run main.go
```

The server will start on `http://localhost:8080`

## 📡 API Endpoints

### WebSocket Connection
- **Endpoint**: `/ws`
- **Protocol**: WebSocket
- **Description**: Establishes a WebSocket connection for real-time messaging

#### Example Connection (JavaScript)
```javascript
const ws = new WebSocket('ws://localhost:8080/ws');

ws.onopen = function() {
    console.log('Connected to chat server');
};

ws.onmessage = function(event) {
    console.log('Received:', event.data);
};

ws.send('Hello, World!');
```

## 🔧 Usage

### Basic WebSocket Communication
The server currently implements a simple echo functionality. When a client sends a message, the server responds with "Echo " + the original message.

### Encryption Implementation
The crypto package provides utilities for implementing end-to-end encryption:

```go
// Generate DH key pair
private, public, err := crypto.GenerateKeyPair()

// Compute shared secret (after receiving peer's public key)
sharedSecret := crypto.SharedSecretKey(private, peerPublic)

// Derive AES key
aesKey := crypto.DerivesAesKey(sharedSecret)

// Encrypt message
ciphertext, err := crypto.Encrypt([]byte("Hello, World!"), aesKey)

// Decrypt message
plaintext, err := crypto.Decrypt(ciphertext, aesKey)
```

## 🚧 Roadmap

- [ ] Complete authentication system
- [ ] Implement message persistence
- [ ] Add user management
- [ ] Room/channel support
- [ ] Message history
- [ ] File sharing capabilities
- [ ] Rate limiting
- [ ] Comprehensive logging

## 🤝 Contributing
I created this project as it would be a good oppurtunity for me to learn about the Diffie–Hellman key exchange algorithm so feel free to wonder around this repo and make pull requests. Let's learn together

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 📞 Support

For questions and support, please open an issue in the GitHub repository.

---

Built with ❤️ using Go and modern cryptographic practices.
