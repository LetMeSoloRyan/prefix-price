# prefix-price

1 Programming Exercise


1.1 Problem

Write a Golang program that allows the user to insert price quotes from some arbitrary numbers (can be more than 2) of companies and to query which company could offer the cheapest price for a specific number. Assume all entries could fit in the memory of an ordinary home computer.

When a number matches several entries of a company, the longest matching prefix of that company should always be used. For example, according to the example price list below, dialing 17841025028 will cost you $2.23/min with Company A (as the number starts with 17) or $2.0/min with Company B (as the number starts with 178).

Some prefixes might not be available in some or even all companies. For example, you can dial 85281777778 only with Company A. On the other hand, the number starting with 84 is not available for both companies.




1.2 Specifications

Your solution should provide two functionalities, INSERT and QUERY via an interactive command line interface. It should be able to read an INSERT command followed by company’s name, phone prefix and price, as well as a QUERY command followed by a phone number.

Your solution should return the query result (the best offer, according to the rules given on page 1) after a QUERY command. If the phone number is not supported by any company, put down NA right after the phone number.

Values should be separated with space for both input and output.

The INSERT command is of the following format:
INSERT [operator] [prefix] [price]

The QUERY input is of the following format:
QUERY [phone no.]

The QUERY output is of the following format:
[phone no.] [operator] [prefix] [best price]

Please see the following sample for the expected input and output formats.



Sample: Program Input & Output


>>> INSERT A 1 0.85
>>> INSERT A 17 2.23 
>>> INSERT A 1787 2.86 
>>> INSERT A 81 0.0 
>>> INSERT B 1 1.0
>>> INSERT B 178 2.0 
>>> INSERT B 44 3.0 
>>> QUERY 17871234567 
17871234567 B 178 2.0
>>> QUERY 810112716677 
810112716677 A 81 0.0 
>>> QUERY 999
999 NA




1.3 Requirements & Constraints

- Use Golang.
- Do not use a database.
- Do not create a GUI.
- Apply the concept of “MVC (Model-View-Controller)”
- Apply the concept of “Dependency injection”
- Apply the concept of “Database/Model normalization”
- The code should be fully unit tested.
- A brief explanation of your design.


1.4 Evaluation Metrics

- Code readability (Important!).
- Code simplicity.
- Correctness.
- Code design.
- Good performance.
