package pubnubTests

import (
    "testing"
    "pubnubMessaging"
    "fmt"
    "encoding/json"
    "unicode/utf16"
)

// Start indicator
func TestEncryptionStart(t *testing.T){
    PrintTestMessage("==========Encryption tests start==========")
}

// TestYayDecryptionBasic tests the yay decryption.
// Assumes that the input message is deserialized  
// Decrypted string should match yay!
func TestYayDecryptionBasic(t *testing.T) {
    message := "q/xJqqN6qbiZMXYmiQC1Fw==";
    //decrypt
    decrypted := pubnubMessaging.DecryptString("enigma", message)

    if("yay!" == decrypted){        
        fmt.Println("Yay decryption basic: passed.") 
    } else {
        t.Error("Yay decryption basic: failed.");
    }
}

// TestYayEncryptionBasic tests the yay encryption.
// Assumes that the input message is not serialized  
// Decrypted string should match q/xJqqN6qbiZMXYmiQC1Fw==
func TestYayEncryptionBasic(t *testing.T) {
    message := "yay!";
    //encrypt
    encrypted := pubnubMessaging.EncryptString("enigma", message);

    if("q/xJqqN6qbiZMXYmiQC1Fw==" == encrypted){
        fmt.Println("Yay encryption basic: passed.") 
    } else {
        t.Error("Yay encryption basic: failed.");
    }
}

// TestYayDecryption tests the yay decryption.
// Assumes that the input message is serialized  
// Decrypted string should match yay!
func TestYayDecryption(t *testing.T) {
    message := "Wi24KS4pcTzvyuGOHubiXg==";
    //decrypt
    decrypted := pubnubMessaging.DecryptString("enigma", message)
    //serialize
    b, err := json.Marshal("yay!")
    if err != nil {
        fmt.Println("error:", err)
        t.Error("Yay decryption: failed.");
    }else {
        if(string(b) == decrypted){        
            fmt.Println("Yay decryption: passed.") 
        } else {
            t.Error("Yay decryption: failed.");
        }
    }    
}

// TestYayEncryption tests the yay encryption.
// Assumes that the input message is serialized  
// Decrypted string should match q/xJqqN6qbiZMXYmiQC1Fw==
func TestYayEncryption(t *testing.T) {
    message := "yay!";
    //serialize
    b, err := json.Marshal(message)
    if err != nil {
        fmt.Println("error:", err)
        t.Error("My object encryption: failed.");
    }else {
        //encrypt
        encrypted := pubnubMessaging.EncryptString("enigma", string(b));
        if("Wi24KS4pcTzvyuGOHubiXg==" == encrypted){
            fmt.Println("Yay encryption: passed.") 
        } else {
            t.Error("Yay encryption: failed.");
        }
    }    
}

// TestArrayDecryption tests the slice decryption.
// Assumes that the input message is deserialized  
// And the output message has to been deserialized.
// Decrypted string should match Ns4TB41JjT2NCXaGLWSPAQ==
func TestArrayDecryption(t *testing.T) {
    message := "Ns4TB41JjT2NCXaGLWSPAQ=="
    //decrypt
    decrypted := pubnubMessaging.DecryptString("enigma", message);
    /*b1, err1 := json.Marshal("string")
    if err1 != nil {
        fmt.Println("error:", err1)
        t.Error("Slice decryption: failed.");
    } else {
        fmt.Println("s:", string(b1))
    }*/
    slice := []string{}
    b, err := json.Marshal(slice)
    if err != nil {
        fmt.Println("error:", err)
        t.Error("Slice decryption: failed.");
    } else {
        if(string(b) == decrypted){        
            fmt.Println("Slice decryption: passed.") 
        } else {
            t.Error("Slice decryption: failed.");
        }
    }    
}

// TestArrayEncryption tests the slice encryption.
// Assumes that the input message is not serialized  
// Decrypted string should match Ns4TB41JjT2NCXaGLWSPAQ==
func TestArrayEncryption(t *testing.T) {
    message := []string{}
    //serialize
    b, err := json.Marshal(message)
    if err != nil {
        fmt.Println("error:", err)
        t.Error("Slice encryption: failed.");
    }else {
        //encrypt
        encrypted := pubnubMessaging.EncryptString("enigma", string(b));
        if("Ns4TB41JjT2NCXaGLWSPAQ==" == encrypted){
            fmt.Println("Slice encryption: passed.") 
        } else {
            t.Error("Slice encryption: failed.");
        }
    }
}

// TestObjectDecryption tests the empty object decryption.
// Assumes that the input message is deserialized  
// And the output message has to been deserialized.
// Decrypted string should match IDjZE9BHSjcX67RddfCYYg==
func TestObjectDecryption(t *testing.T) {
    message := "IDjZE9BHSjcX67RddfCYYg=="
    //decrypt
    decrypted := pubnubMessaging.DecryptString("enigma", message);

    emptyStruct := EmptyStruct{}
    b, err := json.Marshal(emptyStruct)
    if err != nil {
        fmt.Println("error:", err)
        t.Error("Object decryption: failed.");
    } else {
        if(string(b) == decrypted){        
            fmt.Println("Object decryption: passed.") 
        } else {
            t.Error("Object decryption: failed.");
        }
    }    
}

// TestObjectEncryption tests the empty object encryption.
// The output is not serialized
// Encrypted string should match the serialized object
func TestObjectEncryption(t *testing.T) {
    message := EmptyStruct{}
    //serialize
    b, err := json.Marshal(message)
    if err != nil {
        fmt.Println("error:", err)
        t.Error("Object encryption: failed.");
    }else {
        //encrypt
        encrypted := pubnubMessaging.EncryptString("enigma", string(b));
        if("IDjZE9BHSjcX67RddfCYYg==" == encrypted){
            fmt.Println("Object encryption: passed.") 
        } else {
            t.Error("Object encryption: failed.");
        }
    }
}

// TestMyObjectDecryption tests the custom object decryption.
// Assumes that the input message is deserialized  
// And the output message has to been deserialized.
// Decrypted string should match BMhiHh363wsb7kNk7krTtDcey/O6ZcoKDTvVc4yDhZY=
func TestMyObjectDecryption(t *testing.T) {
    message := "BMhiHh363wsb7kNk7krTtDcey/O6ZcoKDTvVc4yDhZY="
    //decrypt
    decrypted := pubnubMessaging.DecryptString("enigma", message);
    customStruct := CustomStruct{
        Foo : "hi!",
        Bar : []int{1,2,3,4,5},
    }
    b, err := json.Marshal(customStruct)
    if err != nil {
        fmt.Println("error:", err)
        t.Error("My object decryption: failed.");
    } else {
        if(string(b) == decrypted){        
            fmt.Println("My object decryption: passed.") 
        } else {
            t.Error("My object decryption: failed.");
        }
    }    
}

// TestMyObjectEncryption tests the custom object encryption.
// The output is not serialized
// Encrypted string should match the serialized object
func TestMyObjectEncryption(t *testing.T) {
    message := CustomStruct{
        Foo: "hi!",
        Bar: []int{1,2,3,4,5},
    }
    //serialize
    b1, err := json.Marshal(message)
    if err != nil {
        fmt.Println("error:", err)
        t.Error("My object encryption: failed.");
    }else {
        //encrypt
        encrypted := pubnubMessaging.EncryptString("enigma", string(b1));
        if("BMhiHh363wsb7kNk7krTtDcey/O6ZcoKDTvVc4yDhZY=" == encrypted){
            fmt.Println("My object encryption: passed.") 
        } else {
            t.Error("My object encryption: failed.");
        }
    }
}

// TestPubNubDecryption2 tests the Pubnub Messaging API 2 decryption.
// Assumes that the input message is deserialized  
// Decrypted string should match Pubnub Messaging API 2
func TestPubNubDecryption2(t *testing.T) {
    message := "f42pIQcWZ9zbTbH8cyLwB/tdvRxjFLOYcBNMVKeHS54=";
    //decrypt
    decrypted := pubnubMessaging.DecryptString("enigma", message)
    b, err := json.Marshal("Pubnub Messaging API 2")
    if err != nil {
        fmt.Println("error:", err)
        t.Error("Pubnub Messaging API 2 decryption: failed.");
    } else {        
        if(string(b) == decrypted){        
            fmt.Println("Pubnub Messaging API 2 decryption: passed.") 
        } else {
            t.Error("Pubnub Messaging API 2 decryption: failed.");
        }
    }    
}

// TestPubNubEncryption2 tests the Pubnub Messaging API 2 encryption.
// Assumes that the input message is not serialized  
// Decrypted string should match f42pIQcWZ9zbTbH8cyLwB/tdvRxjFLOYcBNMVKeHS54=
func TestPubNubEncryption2(t *testing.T) {
    message := "Pubnub Messaging API 2";
    b, err := json.Marshal(message)
    if err != nil {
        fmt.Println("error:", err)
        t.Error("Pubnub Messaging API 2 encryption: failed.");
    } else {    
        //encrypt
        encrypted := pubnubMessaging.EncryptString("enigma", string(b));
        if("f42pIQcWZ9zbTbH8cyLwB/tdvRxjFLOYcBNMVKeHS54=" == encrypted){
            fmt.Println("Pubnub Messaging API 2 encryption: passed.") 
        } else {
            t.Error("Pubnub Messaging API 2 encryption: failed.");
        }
    }        
}

// TestPubNubDecryption tests the Pubnub Messaging API 1 decryption.
// Assumes that the input message is deserialized  
// Decrypted string should match Pubnub Messaging API 1
func TestPubNubDecryption(t *testing.T) {
    message := "f42pIQcWZ9zbTbH8cyLwByD/GsviOE0vcREIEVPARR0=";
    //decrypt
    decrypted := pubnubMessaging.DecryptString("enigma", message)
    
    b, err := json.Marshal("Pubnub Messaging API 1")
    if err != nil {
        fmt.Println("error:", err)
        t.Error("Pubnub Messaging API 1 decryption: failed.");
    } else {    
        if(string(b) == decrypted){        
            fmt.Println("Pubnub Messaging API 1 decryption: passed.") 
        } else {
            t.Error("Pubnub Messaging API 1 decryption: failed.");
        }
    }    
}

// TestPubNubEncryption tests the Pubnub Messaging API 1 encryption.
// Assumes that the input message is not serialized  
// Decrypted string should match f42pIQcWZ9zbTbH8cyLwByD/GsviOE0vcREIEVPARR0=
func TestPubNubEncryption(t *testing.T) {
    message := "Pubnub Messaging API 1";
    b, err := json.Marshal(message)
    if err != nil {
        fmt.Println("error:", err)
        t.Error("Pubnub Messaging API 1 encryption: failed.");
    } else {        
        //encrypt
        encrypted := pubnubMessaging.EncryptString("enigma", string(b));
        if("f42pIQcWZ9zbTbH8cyLwByD/GsviOE0vcREIEVPARR0=" == encrypted){
            fmt.Println("Pubnub Messaging API 1 encryption: passed.") 
        } else {
            t.Error("Pubnub Messaging API 1 encryption: failed.");
        }
    }    
}

// TestStuffCanDecryption tests the StuffCan decryption.
// Assumes that the input message is deserialized  
// Decrypted string should match {\"this stuff\":{\"can get\":\"complicated!\"}}
func TestStuffCanDecryption(t *testing.T) {
    message := "zMqH/RTPlC8yrAZ2UhpEgLKUVzkMI2cikiaVg30AyUu7B6J0FLqCazRzDOmrsFsF";
    //decrypt
    decrypted := pubnubMessaging.DecryptString("enigma", message)
    if("{\"this stuff\":{\"can get\":\"complicated!\"}}" == decrypted){        
        fmt.Println("StuffCan decryption: passed.") 
    } else {
        t.Error("StuffCan decryption: failed.");
    }
}

// TestStuffCanEncryption tests the StuffCan encryption.
// Assumes that the input message is not serialized  
// Decrypted string should match zMqH/RTPlC8yrAZ2UhpEgLKUVzkMI2cikiaVg30AyUu7B6J0FLqCazRzDOmrsFsF
func TestStuffCanEncryption(t *testing.T) {
    message := "{\"this stuff\":{\"can get\":\"complicated!\"}}";
    //encrypt
    encrypted := pubnubMessaging.EncryptString("enigma", message);
    if("zMqH/RTPlC8yrAZ2UhpEgLKUVzkMI2cikiaVg30AyUu7B6J0FLqCazRzDOmrsFsF" == encrypted){
        fmt.Println("StuffCan encryption: passed.") 
    } else {
        t.Error("StuffCan encryption: failed.");
    }
}

// TestHashDecryption tests the hash decryption.
// Assumes that the input message is deserialized  
// Decrypted string should match {\"foo\":{\"bar\":\"foobar\"}}
func TestHashDecryption(t *testing.T) {
    message := "GsvkCYZoYylL5a7/DKhysDjNbwn+BtBtHj2CvzC4Y4g=";
    //decrypt
    decrypted := pubnubMessaging.DecryptString("enigma", message)
    if("{\"foo\":{\"bar\":\"foobar\"}}" == decrypted){        
        fmt.Println("Hash decryption: passed.") 
    } else {
        t.Error("Hash decryption: failed.");
    }
}

// TestHashEncryption tests the hash encryption.
// Assumes that the input message is not serialized  
// Decrypted string should match GsvkCYZoYylL5a7/DKhysDjNbwn+BtBtHj2CvzC4Y4g=
func TestHashEncryption(t *testing.T) {
    message := "{\"foo\":{\"bar\":\"foobar\"}}";
    //encrypt
    encrypted := pubnubMessaging.EncryptString("enigma", message);
    if("GsvkCYZoYylL5a7/DKhysDjNbwn+BtBtHj2CvzC4Y4g=" == encrypted){
        fmt.Println("Hash encryption: passed.") 
    } else {
        t.Error("Hash encryption: failed.");
    }
}

// TestUnicodeDecryption tests the Unicode decryption.
// Assumes that the input message is deserialized  
// Decrypted string should match 漢語
func TestUnicodeDecryption(t *testing.T) {
   message := "+BY5/miAA8aeuhVl4d13Kg==";
    //decrypt
    decrypted := pubnubMessaging.DecryptString("enigma", message)
    data, _, _, err := pubnubMessaging.ParseJson([]byte(decrypted))
    if(err != nil){
        t.Error("Unicode decryption: failed.", err);
    } else {
        if("漢語" == data){        
            fmt.Println("Unicode decryption: passed.") 
        } else {
            t.Error("Unicode decryption: failed.");
        }
    }     
}

// UTF16ToString returns the UTF-8 encoding of the UTF-16 sequence s,
// with a terminating NUL removed.
func UTF16ToString(s []uint16) []rune {
    for i, v := range s {
         if v == 0 {
            s = s[0:i]
            break
        }
    }
    return utf16.Decode(s)
}

// TestUnicodeEncryption tests the Unicode encryption.
// Assumes that the input message is not serialized  
// Decrypted string should match +BY5/miAA8aeuhVl4d13Kg==
func TestUnicodeEncryption(t *testing.T) {
    message := "漢語"
    b, err := json.Marshal(message)
    if err != nil {
        fmt.Println("error:", err)
        t.Error("Unicode encryption: failed.");
    } else {        
        //encrypt
        encrypted := pubnubMessaging.EncryptString("enigma", string(b));
        if("+BY5/miAA8aeuhVl4d13Kg==" == encrypted){
            fmt.Println("Unicode encryption: passed.") 
        } else {
            t.Error("Unicode encryption: failed.");
        }
    }  
}

// TestGermanDecryption tests the German decryption.
// Assumes that the input message is deserialized  
// Decrypted string should match ÜÖ
func TestGermanDecryption(t *testing.T) {
    message := "stpgsG1DZZxb44J7mFNSzg==";
    //decrypt
    decrypted := pubnubMessaging.DecryptString("enigma", message)
    data, _, _, err := pubnubMessaging.ParseJson([]byte(decrypted))
    if(err != nil){
        t.Error("German decryption: failed.", err);
    } else {
        if("ÜÖ" == data){        
            fmt.Println("German decryption: passed.") 
        } else {
            t.Error("German decryption: failed.");
        }
    }
}

// TestGermanEncryption tests the German encryption.
// Assumes that the input message is not serialized  
// Decrypted string should match stpgsG1DZZxb44J7mFNSzg==
func TestGermanEncryption(t *testing.T) {
    message := "ÜÖ"
    b, err := json.Marshal(message)
    if err != nil {
        fmt.Println("error:", err)
        t.Error("Unicode encryption: failed.");
    } else {        
        //encrypt
        encrypted := pubnubMessaging.EncryptString("enigma", string(b));
        if("stpgsG1DZZxb44J7mFNSzg==" == encrypted){
            fmt.Println("German encryption: passed.") 
        } else {
            t.Error("German encryption: failed.");
        }
    }    
}

// End indicator
func TestEncryptionEnd(t *testing.T){
    PrintTestMessage("==========Encryption tests end==========")
}
