package main

import (
        "fmt"
        "strings"
        "cut_dup"
)
var listA = []string{}
var newList = []string{}
var speclistA = []string{}
var newA = []string{}
var noDup = []string{} 


type Target struct{
    
    name string
    surname string
    nick string
    birthdate string
    partner string
    child string
    pet string
    company string
    others string
}

func collectDonnees(){
        var tgt Target
        fmt.Println("MAINTENANT, ENTREZ QUELQUE DONNEES POUR LA CIBLE, SVP:")
        
        fmt.Print("LE PRENOM: ")
        fmt.Scanln(&tgt.name)
        prenom := strings.ToLower(tgt.name)
        
        fmt.Print("LE NOM: ")
        fmt.Scanln(&tgt.surname)
        nom := strings.ToLower(tgt.surname)
        
        fmt.Print("LE NICK, SVP: ")
        fmt.Scanln(&tgt.nick)
        lesurnom := strings.ToLower(tgt.nick)
        
        fmt.Print("LA DATE DE NAISSANCE, SVP: ")
        fmt.Scanln(&tgt.birthdate)
        dn := strings.ToLower(tgt.birthdate)
        
        fmt.Print("LE/A CONSORT(E), SVP: ")
        fmt.Scanln(&tgt.partner)
        consort := strings.ToLower(tgt.partner)
        
        fmt.Print("LE FIL(E), SVP: ")
        fmt.Scanln(&tgt.child)
        fils := strings.ToLower(tgt.child)
        
        fmt.Print("LE CHOUCHOU, SVP: ")
        fmt.Scanln(&tgt.pet)
        lechouchou := strings.ToLower(tgt.pet)
        
        fmt.Print("LA FIRME, SVP: ")
        fmt.Scanln(&tgt.company)
        lafirme := strings.ToLower(tgt.company)
        
        fmt.Print("ET MAINTENANT, METTEZ QUELQUES AUTRES MOT-CLES LIES A LA CIBLE\nEN LES SEPARANT PAR ",".\n>")
        fmt.Scanln(&tgt.others)
        autresmots := strings.ToLower(tgt.others)
        newList = strings.Split(autresmots, ",")
        
        
        listA = append(newList, prenom, nom, lesurnom, dn, consort, fils, lechouchou, lafirme)
        fmt.Println(listA)
        
        
    }
    
    
    
func speChar(list, new []string){
    
    for _, strinGo := range list {
    uni := strings.Replace(strinGo, "a","@", -1)
    duni := strings.Replace(uni, "i","%", -1)
    te := strings.Replace(duni, "h","#", -1)
    salamee := strings.Replace(te, "o","*", -1)
    new = append(new, salamee)
    new = append(new, list[:]...)
    }
    
}

//func randomic(){}

func main(){
    collectDonnees()
    speChar(listA, speclistA)
    fmt.Println(speclistA)
    cut_dup(listA,noDup)
    
    }
