data = ""
with open("./test/Test_sol_Test.bin","r") as f:
	line = f.readlines()
	line = line[-1]
	data = str(line)
    	
with open("./test/Test_sol_Test.bin","wb") as f2:   
	f2.truncate()
	f2.write(data[:-1])
    

with open("./test/Test_sol_Test.abi","r") as f3:  
	line = f3.readlines()
	line = str(line[3:])




with open("./test/Test_sol_Test.abi","w") as f4:  
	f4.truncate()
	f4.write(line[2:-2])
