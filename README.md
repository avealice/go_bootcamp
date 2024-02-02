# go bootcamp

__day00__


The project involves creating a program in the Go programming language that processes a set of integer numbers within the range of -100000 to 100000, read from standard input. The program calculates and displays four statistical metrics: Mean, Median, Mode, and Standard Deviation. Users have the flexibility to choose which specific metrics to print, with the default being all four. The implementation utilizes the standard sorting algorithm for integers in Go and considers both population and regular standard deviation. The goal is to provide a versatile tool for analyzing numerical data sets with customizable output.

__day01__

* readDB: This command-line tool enables the straightforward reading of databases, supporting both XML and JSON formats. It employs the DBReader interface, allowing different implementations for each format while ensuring consistent object types.

* compareDB: Designed for database comparison, compareDB identifies modifications like added or removed cakes, changes in cooking time, and variations in ingredients. It seamlessly works with XML and JSON formats, making it versatile for comparing original and stolen databases.

* compareFS: The compareFS tool facilitates the efficient comparison of server filesystem backups, identifying added and removed files from two plain text file dumps. It handles large file sizes by avoiding simultaneous loading into RAM, offering a practical solution for filesystem backup comparisons.
