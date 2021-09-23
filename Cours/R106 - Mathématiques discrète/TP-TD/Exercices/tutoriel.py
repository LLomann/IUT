import math as mt

def g(x):
	if x<0 :
		return mt.sqrt(-x)
	else :
		return mt.sqrt(x)	
def E():		
	for i in range(30,3,-7) :
		print (i)
	
def h(x):
	m = 10
	for i in range (0,x+1):
		print(m,"x", i," =",m*i )

squares = [1, 4, 9, 16, 25]

def compte(x) :
	liste = [1, 2, 3, 4, 5]
	for i in range (len(liste)) :
		if x == liste[i] :
			print(liste[i])
		else

compte(6)
