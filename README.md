<div align="center">

# Quantum Energy Teleportation

### Reproducible quantum-circuit demonstrations of quantum energy teleportation on IBM Quantum hardware and local simulators

[![DOI](https://img.shields.io/badge/DOI-10.1103%2FPhysRevApplied.20.024051-blue)](https://doi.org/10.1103/PhysRevApplied.20.024051)
[![arXiv](https://img.shields.io/badge/arXiv-2301.02666-b31b1b.svg)](https://arxiv.org/abs/2301.02666)
[![Qiskit](https://img.shields.io/badge/Qiskit-tested%20with%201.4.0-6929C4)](https://www.ibm.com/quantum/qiskit)
[![IBM Quantum](https://img.shields.io/badge/IBM%20Quantum-Hardware%20%2B%20Simulator-052FAD)](https://www.ibm.com/quantum)
[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

<br>

<img src="QET%20slides.gif" width="72%" alt="Quantum Energy Teleportation overview slides">

</div>

---

## Overview

This repository provides complete quantum-circuit implementations for demonstrating **quantum energy teleportation** on IBM Quantum systems.

The code can be run in two modes:

- **Local simulation** — run the circuits on a classical simulator without an IBM Quantum account.
- **Real quantum hardware** — execute the experiment on IBM Quantum devices with an IBM Quantum account and API token.

The default workflow is designed to be accessible: no IBM Quantum credentials are required for local simulation, and the notebooks are written so that the circuit construction, measurement procedure, and energy-estimation logic can be inspected step by step.

---

## What this repository contains

This project accompanies the experimental demonstration reported in:

> K. Ikeda,  
> **“Demonstration of quantum energy teleportation on superconducting quantum hardware,”**  
> *Physical Review Applied* **20**, 024051 (2023).  
> DOI: [10.1103/PhysRevApplied.20.024051](https://doi.org/10.1103/PhysRevApplied.20.024051)  
> arXiv: [2301.02666](https://arxiv.org/abs/2301.02666)

The repository includes:

- A complete notebook for the original QET demonstration.
- Local simulator workflows for laptop execution.
- IBM Quantum hardware execution workflows.
- A 2025 update using modern Qiskit Runtime workflows.
- Error-mitigation demonstrations using **M3**, **dynamical decoupling**, and **Pauli twirling**.
- Slides and documentation for understanding the physics and implementation.

---

## Highlights

- **Quantum Energy Teleportation on IBM Quantum hardware**
- **No IBM Quantum account required for the default simulator run**
- **Circuit-level implementation of Alice’s measurement and Bob’s conditional operation**
- **Energy-injection and energy-extraction estimation**
- **Updated 2025 notebooks with modern error-mitigation techniques**
- **Compatible with Qiskit 1.4.0 in the 2025 update notebooks**
- **Educational notebooks for both quantum-information and quantum-hardware demonstrations**

---

## Quick start

Clone the repository:

```bash
git clone https://github.com/IKEDAKAZUKI/Quantum-Energy-Teleportation.git
cd Quantum-Energy-Teleportation
```

Create and activate a Python environment:

```bash
python -m venv .venv
source .venv/bin/activate
```

For Windows:

```bash
.venv\Scripts\activate
```

Install the recommended packages:

```bash
pip install qiskit==1.4.0 qiskit-aer qiskit-ibm-runtime mthree numpy matplotlib jupyter
```

Launch Jupyter:

```bash
jupyter notebook
```

Then open:

```text
Quantum_Energy_Teleportation.ipynb
```

---

## Running on a local simulator

The default demonstration can be executed on a classical simulator.  
This mode does **not** require an IBM Quantum account or API token.

Recommended starting point:

```text
Quantum_Energy_Teleportation.ipynb
```

This notebook walks through the basic QET protocol, including:

1. Preparation of the ground state.
2. Alice’s local measurement.
3. Bob’s conditional operation.
4. Estimation of injected and teleported energy.

---

## Running on IBM Quantum hardware

To run the experiment on a real IBM Quantum device:

1. Create or sign in to your IBM Quantum account:  
   [https://www.ibm.com/quantum](https://www.ibm.com/quantum)

2. Obtain your IBM Quantum API token.

3. Replace the placeholder token in the relevant notebook or script:

```python
"My_API_Token"
```

with your own token.

> **Important:** never commit your real API token to a public GitHub repository.  
> For production use, store the token in an environment variable or a local configuration file excluded by `.gitignore`.

---

## Latest update 2025

The `Latest update 2025` directory contains an updated implementation using recent Qiskit workflows and error-mitigation techniques.

```text
Latest update 2025/
├── QET.py
├── QET_Experiment_Estimator.ipynb
└── QET_Experiment_M3_Error_Mitigation.ipynb
```

### Included techniques

- **Estimator-based execution**
- **M3 measurement mitigation**
- **Dynamical decoupling**
- **Pauli twirling**
- **Qiskit Runtime workflows**
- **Fake backend testing**

Recommended notebooks:

- [`QET_Experiment_Estimator.ipynb`](Latest%20update%202025/QET_Experiment_Estimator.ipynb)
- [`QET_Experiment_M3_Error_Mitigation.ipynb`](Latest%20update%202025/QET_Experiment_M3_Error_Mitigation.ipynb)

The 2025 update was tested with:

```text
Qiskit 1.4.0
```

---

## Repository structure

```text
.
├── Quantum_Energy_Teleportation.ipynb
├── Latest update 2025/
│   ├── QET.py
│   ├── QET_Experiment_Estimator.ipynb
│   └── QET_Experiment_M3_Error_Mitigation.ipynb
├── Documentation/
├── PRApplied.pdf
├── QET slides.gif
├── QET slides.pdf
├── CITATION.cff
├── LICENSE
└── README.md
```

---

## Documentation and slides

Additional documentation:

> K. Ikeda,  
> **“Quantum Games and Economics through Teleportation”**  
> March 06, 2025.  
> SSRN: [https://ssrn.com/abstract=5168193](https://ssrn.com/abstract=5168193)

Slides:

- [Overview of Quantum Energy Teleportation — SlideShare](https://www.slideshare.net/slideshow/overview-of-quantum-energy-teleportation/287987616)
- [`QET slides.pdf`](QET%20slides.pdf)

---

## References

### Main experimental paper

K. Ikeda,  
**“Demonstration of quantum energy teleportation on superconducting quantum hardware,”**  
*Physical Review Applied* **20**, 024051 (2023).  
DOI: [10.1103/PhysRevApplied.20.024051](https://doi.org/10.1103/PhysRevApplied.20.024051)  
arXiv: [2301.02666](https://arxiv.org/abs/2301.02666)

### IBM Quantum

IBM Quantum website:  
[https://www.ibm.com/quantum](https://www.ibm.com/quantum)

---

## Citation

If you use this repository in your research, please cite:

```bibtex
@article{PhysRevApplied.20.024051,
  title = {Demonstration of Quantum Energy Teleportation on Superconducting Quantum Hardware},
  author = {Ikeda, Kazuki},
  journal = {Phys. Rev. Appl.},
  volume = {20},
  issue = {2},
  pages = {024051},
  numpages = {12},
  year = {2023},
  month = {Aug},
  publisher = {American Physical Society},
  doi = {10.1103/PhysRevApplied.20.024051},
  url = {https://link.aps.org/doi/10.1103/PhysRevApplied.20.024051}
}
```

You can also use the included [`CITATION.cff`](CITATION.cff) file.

---

## License

This project is released under the [MIT License](LICENSE).

---

<div align="center">

**Quantum Energy Teleportation · IBM Quantum · Qiskit · Error Mitigation**

</div>
