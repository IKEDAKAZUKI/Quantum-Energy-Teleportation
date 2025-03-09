import numpy as np
import matplotlib.pyplot as plt 
from qiskit import * #qiskit
from qiskit.visualization import plot_histogram, plot_bloch_multivector, array_to_latex
#from qiskit.extensions import Initialize
from qiskit import QuantumCircuit, QuantumRegister, transpile, assemble#, execute, Aer, BasicAer
from qiskit_aer import Aer
from qiskit.quantum_info import random_statevector
from qiskit.quantum_info import partial_trace, entropy
import qiskit.quantum_info as qi

from qiskit_ibm_runtime import QiskitRuntimeService, Session, SamplerV2 as Sampler, EstimatorV2 as Estimator
from qiskit.transpiler.preset_passmanagers import generate_preset_pass_manager
from qiskit.quantum_info import SparsePauliOp
from qiskit.providers.jobstatus import JobStatus
from qiskit_ibm_runtime.fake_provider import FakeSherbrooke
import mthree as m3


from IPython.display import HTML, display

def minimal_model_ground_state_circuit(k,h):
    qr = QuantumRegister(2)
    cr = ClassicalRegister(2,"alpha")    
    qc = QuantumCircuit(qr, cr)

    #Prepare the ground state
    alpha=-np.arcsin((1/np.sqrt(2))*(np.sqrt(1+h/np.sqrt(h**2+k**2))))
    
    qc.ry(2*alpha,qr[0])
    qc.cx(qr[0],qr[1])

    return qc

def sin(k,h):
    return (h*k)/np.sqrt((h**2+2*k**2)**2+(h*k)**2)

def inject_energy_circuit(k,h):
    
    qr = QuantumRegister(2)
    cr = ClassicalRegister(2,"alpha")
    
    qc = QuantumCircuit(qr, cr)
    
    #Prepare the ground state
    alpha=-np.arcsin((1/np.sqrt(2))*(np.sqrt(1+h/np.sqrt(h**2+k**2))))
    
    qc.ry(2*alpha,qr[0])
    qc.cx(qr[0],qr[1])
    
    # Alice's projective measurement
    qc.h(qr[0])

    qc.measure(range(2), range(2))

    return qc

def inject_energy_val(k,h,counts_inject,shots):
    ene_A=(h**2)/(np.sqrt(h**2+k**2))
    
    error_A=[]
    
    for orig_bit_string, count in counts_inject.items():
            bit_string = orig_bit_string[::-1]
            ene_A += h*(-1)**int(bit_string[0])*count/shots
            
            for i in range(count):
                error_A.append(h*(-1)**int(bit_string[0]))
            
    print("Injected energy using backend",ene_A,"STD is",np.std(error_A)/np.sqrt(shots))
    print("Exact injected energy",h**2/np.sqrt(h**2+k**2))

def QET_circuit_XX(k,h):
    qr = QuantumRegister(2)
    cr = ClassicalRegister(2,"alpha")
    qc = QuantumCircuit(qr, cr)

    #Prepare the ground state
    alpha=-np.arcsin((1/np.sqrt(2))*(np.sqrt(1+h/np.sqrt(h**2+k**2))))
    
    qc.ry(2*alpha,qr[0])
    qc.cx(qr[0],qr[1])
    
    # Alice's projective measurement
    qc.h(qr[0])
    
    #Bob's conditional operation
    phi=0.5*np.arcsin(sin(k,h))
    qc.cry(-2*phi,qr[0],qr[1])
    
    qc.x(qr[0])
    qc.cry(2*phi,qr[0],qr[1])
    qc.x(qr[0])
    
    #Measurement of the interaction XX
    qc.h(qr[1])

    qc.measure(range(2), range(2))
    
    return qc


def QET_circuit_Z(k,h):
    qr = QuantumRegister(2)
    cr = ClassicalRegister(2,"alpha")
    qc = QuantumCircuit(qr, cr)

    #Prepare the ground state
    alpha=-np.arcsin((1/np.sqrt(2))*(np.sqrt(1+h/np.sqrt(h**2+k**2))))
    
    qc.ry(2*alpha,qr[0])
    qc.cx(qr[0],qr[1])
    
    # Alice's projective measurement
    qc.h(qr[0])
    
    #Bob's conditional operation
    phi=0.5*np.arcsin(sin(k,h))
    qc.cry(-2*phi,qr[0],qr[1])
    
    qc.x(qr[0])
    qc.cry(2*phi,qr[0],qr[1])
    qc.x(qr[0])

    qc.measure(range(2), range(2))
    
    return qc

def QET_energy_XX(k,h,counts_XX,shots):
    ene_XX=(2*k**2)/(np.sqrt(h**2+k**2))
    error_XX=[]
    for orig_bit_string, count in counts_XX.items():
            bit_string = orig_bit_string[::-1]
        
            ene_XX += 2*k*(-1)**int(bit_string[0])*(-1)**int(bit_string[1])*count/shots
    
            for i in range(count):
                error_XX.append(2*k*(-1)**int(bit_string[0])*(-1)**int(bit_string[1]))

    print("Teleported XX energy is",ene_XX,"STD is",np.std(error_XX)/np.sqrt(shots))
    return ene_XX, np.std(error_XX)/np.sqrt(shots)
    
def QET_energy_Z(k,h,counts_Z,shots):
    ene_Z=(h**2)/(np.sqrt(h**2+k**2))
    error_Z=[]

    for orig_bit_string, count in counts_Z.items():
            bit_string = orig_bit_string[::-1]
        
            ene_Z += h*(-1)**int(bit_string[1])*count/shots
            
            for i in range(count):
                error_Z.append(h*(-1)**int(bit_string[1]))
                
    print("Teleported Z energy is",ene_Z,"STD is",np.std(error_Z)/np.sqrt(shots))
    return ene_Z, np.std(error_Z)/np.sqrt(shots)

def QET_QST_XX(k,h):
    qr1 = QuantumRegister(4)
    cr1 = ClassicalRegister(4,"alpha")
    qc1 = QuantumCircuit(qr1, cr1)
    
    #Prepare the ground state
    alpha=-np.arcsin((1/np.sqrt(2))*(np.sqrt(1+h/np.sqrt(h**2+k**2))))
    
    qc1.ry(2*alpha,qr1[0])
    qc1.cx(qr1[0],qr1[1])
    
    # Alice's projective measurement
    qc1.h(qr1[0])

    #Bob's conditional operation
    phi=0.5*np.arcsin(sin(k,h))
    qc1.cry(-2*phi,qr1[0],qr1[1])
    
    qc1.x(qr1[0])
    qc1.cry(2*phi,qr1[0],qr1[1])
    qc1.x(qr1[0])
    
    #Measurement of the interaction XX
    # Comment out qc.h(qr[1]) below for the measurement of Bob's Z term
    qc1.h(qr1[1])
    qc1.measure([1],[1])
    
    # Teleport Quantum State after the measurement 
    #create Bell pair
    qc1.h(qr1[2])
    qc1.cx(qr1[2],qr1[3])
    
    #Bell measurement
    qc1.cx(qr1[1],qr1[2])
    qc1.h(qr1[1])
    
    #Equivalent to conditional operation
    qc1.cx(qr1[2],qr1[3])
    qc1.cz(qr1[1],qr1[3])

    qc1.measure([0,2,3], [0,2,3])
    return qc1

def QET_QST_Z(k,h):
    qr1 = QuantumRegister(4)
    cr1 = ClassicalRegister(4,"alpha")
    qc1 = QuantumCircuit(qr1, cr1)
    
    #Prepare the ground state
    alpha=-np.arcsin((1/np.sqrt(2))*(np.sqrt(1+h/np.sqrt(h**2+k**2))))
    
    qc1.ry(2*alpha,qr1[0])
    qc1.cx(qr1[0],qr1[1])
    
    # Alice's projective measurement
    qc1.h(qr1[0])

    #Bob's conditional operation
    phi=0.5*np.arcsin(sin(k,h))
    qc1.cry(-2*phi,qr1[0],qr1[1])
    
    qc1.x(qr1[0])
    qc1.cry(2*phi,qr1[0],qr1[1])
    qc1.x(qr1[0])

    qc1.measure([1],[1])
    
    # Teleport Quantum State after the measurement 
    #create Bell pair
    qc1.h(qr1[2])
    qc1.cx(qr1[2],qr1[3])
    
    #Bell measurement
    qc1.cx(qr1[1],qr1[2])
    qc1.h(qr1[1])
    
    #Equivalent to conditional operation
    qc1.cx(qr1[2],qr1[3])
    qc1.cz(qr1[1],qr1[3])

    qc1.measure([0,2,3], [0,2,3])
    return qc1

def Confirm_XX_val(k,h,counts_XX,shots):
    ene_XX=(2*k**2)/(np.sqrt(h**2+k**2))

    error_XX=[]
    for orig_bit_string, count in counts_XX.items():
            bit_string = orig_bit_string[::-1]
        
            ene_XX += 2*k*(-1)**int(bit_string[0])*(-1)**int(bit_string[3])*count/shots
    
            for i in range(count):
                error_XX.append(2*k*(-1)**int(bit_string[0])*(-1)**int(bit_string[3]))
    
    print("Confirmed XX energy is",ene_XX,"STD is",np.std(error_XX)/np.sqrt(shots))
    
    return ene_XX, np.std(error_XX)/np.sqrt(shots)

def Confirm_Z_val(k,h,counts_Z,shots):
    ene_Z=(h**2)/(np.sqrt(h**2+k**2))
    error_Z=[]

    for orig_bit_string, count in counts_Z.items():
            bit_string = orig_bit_string[::-1]
        
            ene_Z += h*(-1)**int(bit_string[3])*count/shots
            
            for i in range(count):
                error_Z.append(h*(-1)**int(bit_string[3]))
                
    print("Confirmed Z energy is",ene_Z,"STD is",np.std(error_Z)/np.sqrt(shots))
    return ene_Z, np.std(error_Z)/np.sqrt(shots)

def Confirm_NO_QET_XX(k,h):
    qr = QuantumRegister(4)
    cr = ClassicalRegister(4,"alpha")
    
    qc = QuantumCircuit(qr, cr)
    
    
    #Prepare the ground state
    alpha=-np.arcsin((1/np.sqrt(2))*(np.sqrt(1+h/np.sqrt(h**2+k**2))))
    
    qc.ry(2*alpha,qr[0])
    qc.cx(qr[0],qr[1])
    
    # Alice's projective measurement
    qc.h(qr[0])
    
    
    # Teleport Quantum State 
    #create Bell pair
    qc.h(qr[2])
    qc.cx(qr[2],qr[3])
    
    #Bell measurement
    qc.cx(qr[1],qr[2])
    qc.h(qr[1])
    
    
    #Equivalent to conditional operation
    qc.cx(qr[2],qr[3])
    qc.cz(qr[1],qr[3])

    # Compute XX
    qc.h(qr[3])
    
    qc.measure([0,1,2,3], [0,1,2,3])
    return qc

def Confirm_NO_QET_Z(k,h):
    qr = QuantumRegister(4)
    cr = ClassicalRegister(4,"alpha")
    
    qc = QuantumCircuit(qr, cr)
    
    
    #Prepare the ground state
    alpha=-np.arcsin((1/np.sqrt(2))*(np.sqrt(1+h/np.sqrt(h**2+k**2))))
    
    qc.ry(2*alpha,qr[0])
    qc.cx(qr[0],qr[1])
    
    # Alice's projective measurement
    qc.h(qr[0])
    
    
    # Teleport Quantum State 
    #create Bell pair
    qc.h(qr[2])
    qc.cx(qr[2],qr[3])
    
    #Bell measurement
    qc.cx(qr[1],qr[2])
    qc.h(qr[1])
    
    
    #Equivalent to conditional operation
    qc.cx(qr[2],qr[3])
    qc.cz(qr[1],qr[3])
    
    qc.measure([0,1,2,3], [0,1,2,3])
    return qc

def M3_QEM_Injected(k,h,quasis):
    ene_A=(h**2)/(np.sqrt(h**2+k**2))
    return ene_A+h*quasis.expval("IZ"), quasis.expval_and_stddev("IZ")[1]

def M3_QEM_XX(k,h,quasis):
    ene_XX=(2*k**2)/(np.sqrt(h**2+k**2))
    return ene_XX+2*k*quasis.expval(), quasis.expval_and_stddev()[1]

def M3_QEM_Z(k,h,quasis):
    ene_Z=(h**2)/(np.sqrt(h**2+k**2))
    return ene_Z+h*quasis.expval("ZI"), quasis.expval_and_stddev("ZI")[1]

def M3_QEM_Confirm_XX(k,h,quasis):
    ene_XX=(2*k**2)/(np.sqrt(h**2+k**2))
    return ene_XX+2*k*quasis.expval("ZIIZ"), quasis.expval_and_stddev("ZIIZ")[1]

def M3_QEM_Confirm_Z(k,h,quasis):
    ene_Z=(h**2)/(np.sqrt(h**2+k**2))
    return ene_Z+h*quasis.expval("ZIII"), quasis.expval_and_stddev("ZIII")[1]


def QET_Estimator(k,h):
    qr = QuantumRegister(2)
    cr = ClassicalRegister(2)
    qc = QuantumCircuit(qr, cr)

    #Prepare the ground state
    alpha=-np.arcsin((1/np.sqrt(2))*(np.sqrt(1+h/np.sqrt(h**2+k**2))))
    
    qc.ry(2*alpha,qr[0])
    qc.cx(qr[0],qr[1])
    
    # Alice's projective measurement
    qc.h(qr[0])
    
    #Bob's conditional operation
    phi=0.5*np.arcsin(sin(k,h))
    qc.cry(-2*phi,qr[0],qr[1])
    
    qc.x(qr[0])
    qc.cry(2*phi,qr[0],qr[1])
    qc.x(qr[0])
    
    return qc

def QET_QST_Estimator(k,h):
    qr1 = QuantumRegister(4)
    cr1 = ClassicalRegister(4)
    qc1 = QuantumCircuit(qr1, cr1)
    
    #Prepare the ground state
    alpha=-np.arcsin((1/np.sqrt(2))*(np.sqrt(1+h/np.sqrt(h**2+k**2))))
    
    qc1.ry(2*alpha,qr1[0])
    qc1.cx(qr1[0],qr1[1])
    
    # Alice's projective measurement
    qc1.h(qr1[0])

    #Bob's conditional operation
    phi=0.5*np.arcsin(sin(k,h))
    qc1.cry(-2*phi,qr1[0],qr1[1])
    
    qc1.x(qr1[0])
    qc1.cry(2*phi,qr1[0],qr1[1])
    qc1.x(qr1[0])
    
    
    # Teleport Quantum State after the measurement 
    #create Bell pair
    qc1.h(qr1[2])
    qc1.cx(qr1[2],qr1[3])
    
    #Bell measurement
    qc1.cx(qr1[1],qr1[2])
    qc1.h(qr1[1])
    
    #Equivalent to conditional operation
    qc1.cx(qr1[2],qr1[3])
    qc1.cz(qr1[1],qr1[3])
    
    return qc1

def Confirm_NO_QET_Estimator(k,h):
    qr = QuantumRegister(4)
    cr = ClassicalRegister(4)
    
    qc = QuantumCircuit(qr, cr)
    
    
    #Prepare the ground state
    alpha=-np.arcsin((1/np.sqrt(2))*(np.sqrt(1+h/np.sqrt(h**2+k**2))))
    
    qc.ry(2*alpha,qr[0])
    qc.cx(qr[0],qr[1])
    
    # Alice's projective measurement
    qc.h(qr[0])
    
    
    # Teleport Quantum State 
    #create Bell pair
    qc.h(qr[2])
    qc.cx(qr[2],qr[3])
    
    #Bell measurement
    qc.cx(qr[1],qr[2])
    qc.h(qr[1])
    
    
    #Equivalent to conditional operation
    qc.cx(qr[2],qr[3])
    qc.cz(qr[1],qr[3])

    return qc