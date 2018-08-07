

// ------------------------------------------- Packages ------------------------------------------- //
// IMPORTANT:
// --> Capitalization of the first letter determines access.
// MyFunc() is a public funcion, myFunc() is private.
// 


// ------------------------------------------- Classes Basics ------------------------------------------- //
// No classes, use struct instead
// Structs can have methods

// Good practice:
// Make the struct private (lower case name) and define a New() function, so that struct is properly initialized upon creation 

// ex:

type employee struct {  
    firstName   string
    lastName    string
    totalLeaves int
    leavesTaken int
}

func New(firstName string, lastName string, totalLeave int, leavesTaken int) employee {  
    e := employee {firstName, lastName, totalLeave, leavesTaken}
    return e
}

func (e employee) LeavesRemaining() {  
    fmt.Printf("%s %s has %d leaves remaining", e.firstName, e.lastName, (e.totalLeaves - e.leavesTaken))
}

// ^Call this using:
e := employee.New("Sam", "Adolf", 30, 20)
e.LeavesRemaining()


// ------------------------------------------- Classes Structs ------------------------------------------- //
// No struct "inheritence". Instead use "composition"
// To use methods from a "parent" struct, embed the parent object as a field in the struct
// The child struct can access methods of the parent struct as if it were its own

// ex:


// "Parent" struct
type author struct {  
    firstName string
    lastName  string
    bio       string
}

func (a author) fullName() string {  
    return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

// "Child" struct
type post struct {  
    title     string
    content   string
    author
}

// Now post structs have access to the fullName() method
// ex:
exampleAuthor = Author....
p := post("postTitle", "postContent", exampleAuthor)
p.fullName()


// ------------------------------------------- Classes Interfaces ------------------------------------------- //
// No "inheritence" in Go. Instead use interfaces
// Interface is a type. It's definition is a list of method names.
// If any struct has method implementations for all of the method names in the interface, it can be considered used as the Interface type

// ex:
// Employee is the interface, and Permanent and Contract are two structs that employ the interface.

type Employee interface {  
    CalculateSalary() int
}

type Permanent struct {  
    empId    int
    basicpay int
    pf       int
}

type Contract struct {  
    empId  int
    basicpay int
}

//salary of permanent employee is sum of basic pay and pf
func (p Permanent) CalculateSalary() int {  
    return p.basicpay + p.pf
}

//salary of contract employee is the basic pay alone
func (c Contract) CalculateSalary() int {  
    return c.basicpay
}

/*
total expense is calculated by iterating though the Employee slice and summing  
the salaries of the individual employees  
*/
func totalExpense(s []Employee) {  
    expense := 0
    for _, v := range s {
        expense = expense + v.CalculateSalary()
    }
    fmt.Printf("Total Expense Per Month $%d", expense)
}

func main() {  
    pemp1 := Permanent{1, 5000, 20}
    pemp2 := Permanent{2, 6000, 30}
    cemp1 := Contract{3, 3000}
    employees := []Employee{pemp1, pemp2, cemp1}
    totalExpense(employees)

}



// ------------------------------------------- Loops ------------------------------------------- //

// Standard loop
for i := 0; i <= 10; { // initialisation and post are omitted
	fmt.Printf("%d ", i)
	i += 2
}

// Infinite Loops
for {
}

// Iterate through collection
for index, element := range s {
	fmt.Println(element)
}

// ------------------------------------------- Switch Statements ------------------------------------------- //
switch finger := 8; finger {
case 1:
	fmt.Println("Thumb")
case 2:
	fmt.Println("Index")
case 3:
	fmt.Println("Middle")
case 4:
	fmt.Println("Ring")
case 5:
	fmt.Println("Pinky")
default: //default case
	fmt.Println("incorrect finger number")
}

// Switch Statements - Multiple cases evaluate the same
letter := "i"
switch letter {
case "a", "e", "i", "o", "u": //multiple expressions in case
	fmt.Println("vowel")
default:
	fmt.Println("not a vowel")
}

// Switch Statements - Fallthrough
switch num := number(); { //num is not a constant
case num < 50:
	fmt.Printf("%d is lesser than 50\n", num)
	fallthrough
case num < 100:
	fmt.Printf("%d is lesser than 100\n", num)
	fallthrough
case num < 200:
	fmt.Printf("%d is lesser than 200", num)
}

// ------------------------------------------- Arrays ------------------------------------------- //
// Assignment and passing in by parameter uses a copy, not a reference to the original, EXCEPT FOR SLICES
// Modifying slices MODIFIES ORIGINAL DATA
// ex:
var myArr [3]int
var myArr [3][3]int
myArr := [3]int{12}        // produces [12, 0, 0]
myArr := [...]int{1, 2, 3} // produces [1, 2, 3]
//
// Iterating:
a := [...]float64{67.7, 89.8, 21, 78}
for i, v := range a { //range returns both the index and value
}