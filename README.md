# Calculator

## Goals

This project is a command-line interface (CLI) calculator application that can perform basic arithmetic operations with more than two numbers. The application allow users to enter multiple numbers and select the desired operation (addition, subtraction, multiplication or division) to obtain the result.

## Application assumptions

1. Input Validation
    - The user is only able to input numerical values for the operands.
    - The input can't be left empty.
    - Invalid inputs (such as non-numeric characters) are detected and handled appropriately.
2. Division by Zero
    - If the user enters a division operation with "0" as the divisor, the application display an error message indicating that division by zero is not allowed.
3. Result Display
    - The result of the calculation is displayed in a clear and readable format.
    - If the result is a decimal number, it is rounded to a reasonable number of decimal places.
4. Command Line Interface
    - The CLI have a clear and intuitive syntax for entering commands, operands, and operations.
    - The instructions and usage guidelines are displayed when the application is run.
5. Error Handling
    - The application handle unexpected errors gracefully and display appropriate error messages when necessary.
    - Error messages provide clear instructions on how to correct the issue.
6. Calculation Accuracy
    - The calculator perform accurate calculations, ensuring that the arithmetic operations are executed correctly.
    - Care is taken to handle edge cases and potential rounding errors.
7. Extended Operations
    - The calculator support addition, subtraction, and multiplication with more than two numbers.
    - Users is able to input multiple numbers separated by spaces.
    - The application correctly interpret the order of operations and perform calculations accordingly.
