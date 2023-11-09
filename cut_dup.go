package cut_dup

//import packages

import(
        "fmt"
)

//l'execution de programme, commence ici
func cut_dup(listin, listout []string) []string {
    //declarer une slice avec les valeur en double
    nums_with_dup := []string{}
    
    //declarer un map vide
    mp := make(map[string]bool) 
    
    //declarer une slice vide pour stocker les valeur en double
    nums_no_dup := []string{}   
    
    //parcourir chaque element dans une tranche donnee
    for _, element := range nums_with_dup{
        //verifier si l'element actuel est present sur·le·map
        
        if _, value := mp[element]; !value{
            mp[element] = true
            
            nums_no_dup = append(nums_no_dup, element)
            }
        }
        fmt.Println(nums_no_dup)
        return nums_no_dup
}
