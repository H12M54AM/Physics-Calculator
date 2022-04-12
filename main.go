package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	var choice int

	fmt.Println("\nHello! Please choose the following Calculator!")
	fmt.Println("\nWe have:\n\t1. vf = vi + at\n\t2. d = ((vi + vf) / 2) * t\n\t3. d = vi*t + 0.5 * a * t^2\n\t4. vf^2 = vi^2 + 2 * a * d\n\t5. Work = Force * Displacement\n\t6. eff = (output / input) * 100\n\t7. Ei + w = Ef\n\t8. Torque = F * r")
	fmt.Print("\nWhat would you like to use? : ")
	fmt.Scan(&choice)

	switch choice {
	case 1: // vf = vi + at
		var vf float64
		var vi float64
		var a float64
		var t float64

		fmt.Print("What is the Initial Velocity? (in m/s): ")
		fmt.Scan(&vi)

		fmt.Print("What is the Acceleration? (in m/s^2): ")
		fmt.Scan(&a)

		fmt.Print("What is the time? (in seconds): ")
		fmt.Scan(&t)

		vf = vi + (a * t)

		fmt.Println("\nThe Final Velocity is", vf, "m/s")
		fmt.Println("\nMade by Edward Naidoo - H12M54AM | Github")

	case 2: // d = ((vi + vf) / 2) * t
		var d float64
		var vf float64
		var vi float64
		var t float64

		fmt.Print("What is the Final Velocity? (in m/s): ")
		fmt.Scan(&vf)

		fmt.Print("What is the Initial Velocity? (in m/s): ")
		fmt.Scan(&vi)

		fmt.Print("What is the time? (in seconds): ")
		fmt.Scan(&t)

		d = ((vi + vf) / 2) * t
		fmt.Println("\nThe Displacment is", d, "m")
		fmt.Println("\nMade by Edward Naidoo - H12M54AM | Github")

	case 3: // d = vi*t + 0.5 * a * t^2
		var d float64
		var a float64
		var vi float64
		var t float64

		fmt.Print("What is the Initial Velocity? (in m/s): ")
		fmt.Scan(&vi)

		fmt.Print("What is the time? (in seconds): ")
		fmt.Scan(&t)

		fmt.Print("What is the Accelereation? (in m/s^2): ")
		fmt.Scan(&a)

		d = (vi * t) + 0.5*a*math.Pow(t, 2)
		fmt.Println("The Displacement is", d, "m")
		fmt.Println("\nMade by Edward Naidoo - H12M54AM | Github")

	case 4: // vf^2 = vi^2 + 2 * a * d
		var d float64
		var a float64
		var vi float64
		var vf float64

		fmt.Print("What is the Initial Velocity? (in m/s): ")
		fmt.Scan(&vi)

		fmt.Print("What is the Acceleration? (in m/s^2): ")
		fmt.Scan(&a)

		fmt.Print("What is the Displacment? (in m): ")
		fmt.Scan(&d)

		vf = (math.Pow(vi, 2)) + (2 * a * d)
		vf = math.Sqrt(vf)

		fmt.Println("The Final Velocity is :", vf, "m/s")
		fmt.Println("\nMade by Edward Naidoo - H12M54AM | Github")

	case 5: // Work = Force * Displacement
		var work float64
		var force float64
		var displacement float64

		fmt.Print("What is the Force? (in N): ")
		fmt.Scan(&force)

		fmt.Print("What is the displacement? (in m): ")
		fmt.Scan(&displacement)

		work = force * displacement

		fmt.Println("The work is", work, "N")
		fmt.Println("\nMade by Edward Naidoo - H12M54AM | Github")

	case 6: // eff = (output / input) * 100
		var output float64
		var input float64
		var eff float64

		fmt.Print("What is the Output? : ")
		fmt.Scan(&output)

		fmt.Print("What is the Input? : ")
		fmt.Scan(&input)

		eff = (output / input) * 100
		fmt.Println("The Effeciency is", eff, "%")
		fmt.Println("\nMade by Edward Naidoo - H12M54AM | Github")

	case 7: // Ei + w = Ef
		var newChoice int

		fmt.Print("Which part of Ei + W = Ef would you like to solve for? \n\t1. Ei\n\t2. w\n\t3. Ef\n\n======> ")
		fmt.Scan(&newChoice)

		switch newChoice {
		case 1: 
			solveForEI()
		case 2: 
			solveForW()
		case 3: 
			solveForEF()
		default:
			fmt.Println("Please enter one of the provided Options\n\t1. Ei\n\t2. w\n\t3. Ef")
			fmt.Println("\nMade by Edward Naidoo - H12M54AM | Github")

		}

	case 8: //Tau = Force * Length
		var tau float64 // Torque (in Nm)
		var f float64 // Force (in N)
		var r float64 // Radius from Pivot (in m)
		
		fmt.Print("What is the Force? => ")
		fmt.Scan(&f)

		fmt.Print("What is the Length or Radius from the Pivot? => ")
		fmt.Scan(&r)

		tau = f * r

		fmt.Println("\nYou got", tau, "Nm of Torque")
		fmt.Println("\nMade by Edward Naidoo - H12M54AM | Github")

	default:
		fmt.Println("Please enter the chosen number..")
	}
}

//---------------------Functions------------------------------

func solveForEI() {
	var response string
	var W float64
	var Ef, Ei float64

	// Work
	fmt.Print("Does this problem already have Work? => ")
	fmt.Scan(&response)

	if response == "yes" || response != "no" {
		fmt.Print("What is the Work value? => ")
		fmt.Scan(&W)
		
	} else {
		W += 0
	}

	// Ef
	fmt.Print("Does this problem already have Final Energy? => ")
	fmt.Scan(&response)

	var lower = strings.ToLower(response)

	if lower == "yes" || lower != "no" {
		fmt.Print("What is the Final Energy? => ")
		fmt.Scan(&Ef)
		
	} else {
		Ef += 0
	}

	Ei = Ef - W
	fmt.Println("Energy Initial is", Ei, "Joules")
	fmt.Println("\nMade by Edward Naidoo - H12M54AM | Github")

}

func solveForW() {
	var response string
	var W float64
	var Ef, Ei float64

	// Ei
	fmt.Print("Does this problem already have Ei? => ")
	fmt.Scan(&response)
	var lower = strings.ToLower(response)

	if lower == "yes" || lower != "no" {
		fmt.Print("What is Ei? => ")
		fmt.Scan(&Ei)
		
	} else {
		Ei += 0
	}

	// Ef
	fmt.Print("Does this problem already have Final Energy? => ")
	fmt.Scan(&response)

	if lower == "yes" || lower != "no" {
		fmt.Print("What is the Final Energy? => ")
		fmt.Scan(&Ef)
		
	} else {
		Ef += 0
	}

	W = Ef - Ei
	fmt.Println("Work is", W, "Joules")
	fmt.Println("\nMade by Edward Naidoo - H12M54AM | Github")

}

func solveForEF() {
	var response string
	var W float64
	var Ef, Ei float64
	

	// Ei
	fmt.Print("Does this problem already have Ei? => ")
	fmt.Scan(&response)
	var lower = strings.ToLower(response)

	if lower == "yes" || lower != "no" {
		fmt.Print("What is Ei? => ")
		fmt.Scan(&Ei)
		
	} else {
		Ei += 0
	}

	//Work
	fmt.Print("Does this problem already have Work? => ")
	fmt.Scan(&response)

	if lower == "yes" || lower != "no" {
		fmt.Print("What is the Work value? => ")
		fmt.Scan(&W)
		
	} else {
		W += 0
	}

	Ef = Ei + W
	fmt.Println("Energy Final is", Ef, "Joules")
	fmt.Println("\nMade by Edward Naidoo - H12M54AM | Github")

}