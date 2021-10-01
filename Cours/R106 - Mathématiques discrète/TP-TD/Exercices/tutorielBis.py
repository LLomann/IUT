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

def compte(l):
	compteur = 0
	for compteur in range (0,len(l)):
		compteur = compteur + 1
	return compteur
		

# Exercice 4.1

def et(x,y):
    if x:
        return bool(y)
    else:
        return False
#teste
#print(et(1,0))

def etBIS(x, y):
	if x:
		if y:
			return True
		else:
			return False
	else:
		return False


def ou(x,y):
    if x:
        return x
    else:
        return y

#teste
#print(ou(0,1))

def non(x):
    if x == True:
        return False
    else:
        return True

#teste
#print(non(1))


#Exercice 5
#a)
def affiche_col(t):
    print("Ensemble:{")
    for e in t:
        print(e)
    print("}")

#teste
#print(affiche_col(set())) #enesemble vide
#print(affiche_col({1,2,5,8}))
a = set()
a.add(3)
a.add(5)
a.add(8)
#print(affiche_col(a))

#b)
def membre(x,E):
    for i in E:
        if i == x:
            return True
    else:
        return False
    
#teste
#print(membre(8, {1,2,5,3,6,7,8,4}))

#c)
def somme(E):
    somme = 0
    for i in E:
        somme = somme + i
    return somme

#teste
#print(somme({3,1,2})) #ne fonctionne pas avec des nombre double :/
#print(somme({1,1}))

#d)
def taille(E): 
    compte = 0
    for i in E:
        compte = compte+1
    return compte
    

#teste
#print(taille({1,3,4,5})) #ne marche pas pour les doublons :/

#e)
def moyenne(E):
    somme = 0
    compte = 0
    for i in E:
        somme = somme + i
        compte = compte+1
    return somme/compte

#teste
#print(moyenne({10,14,16,20}))

# Exercice 5.2
#a)
def est_ensemble(l) :
    for i in range(len(l)) :
        for j in range(i+1, len(l)) :
            if l[i] == l[j] :   
                return False
            j = j+1
        i = i+1
    return True

liste = [1,2,3,2,4,8,9,7]
print(est_ensemble(liste))
		
#b)
def suppr_doublon(l) :
    tab = []
    for i in range(len(l)) :
        double = False
        for j in range(i+1, len(l)) :
            if l[i] == l[j] :
                double = True
                break
        if not double : 
            tab.append(l[i])         
    return tab        
                       
print(suppr_doublon(liste))

