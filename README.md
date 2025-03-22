# ResQube 🚀

ResQube is a CLI tool that recommends **initial CPU and memory requests & limits** for Kubernetes services based on their **language, framework, and workload type**. It helps developers and SRE teams set optimal resource configurations to improve **performance and cost-efficiency**.

## 📌 Features
- ✅ Supports **Java, Python, JavaScript**
- ✅ Recommends **CPU/memory requests & limits**
- ✅ Generates **Kubernetes Deployment YAML**
- ✅ Extensible for **frameworks & custom tuning**
- ✅ Optional future integration with **Prometheus metrics**

## 🛠 Installation

Make sure you have **Go installed** (version 1.18+):

```sh
# Clone the repository
git clone https://github.com/yourusername/resqube.git
cd resqube

# Build the binary
go build -o resqube

# Run the CLI
./resqube --help
```

## 🚀 Usage

Run the CLI with the required parameters:

```sh
./resqube --language=java --workload=medium
```

## 🔧 Development & Contribution

1. Fork the repository and clone it locally.
2. Install dependencies and build the project.
3. Submit a pull request with improvements!

## 📜 License

ResQube is released under the **MIT License**.

---
💡 **Future Enhancements:**
- 📊 **Prometheus Integration** (real-time resource usage)
- ⚙ **Support for More Languages & Frameworks**
- 🌍 **Helm Chart Generation**

🚀 **Contributions Welcome!** 🎉
