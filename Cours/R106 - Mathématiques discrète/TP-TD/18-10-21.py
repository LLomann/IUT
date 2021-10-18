###################################################################################################################################################################

def reste(a,b):
    while a >= b :
        a = a-b
    return a

#test
assert reste(1,5)==1
assert reste(21,5)==1
assert reste(32,6)==2
assert reste(12,4)==0
assert reste(12,9)==3

###################################################################################################################################################################

def divisible(a,b):
    return reste(a,b) == 0

#test
assert divisible(12,4)
assert not divisible(32,6)

###################################################################################################################################################################

def diviseur(a) :
    liste=[]
    for i in range (2, int (a/2)+1) :
        if divisible(a,i) :
            liste.append(i)
    return liste

#test
assert diviseur(6) == [2,3]
print(diviseur(10000))

###################################################################################################################################################################

def premier(a):
    return diviseur(a) == [] 

#test
assert not premier(6)
assert premier(11)

###################################################################################################################################################################

def tout_les_premier(a):
    liste =[]
    for i in range (0,a) :
        if premier (i):
            liste.append(i)
    return liste

#test
assert tout_les_premier(13) == [0, 1, 2, 3, 5, 7, 11]
print( tout_les_premier(7000))

###################################################################################################################################################################