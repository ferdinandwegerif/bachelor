

window.vuguRender = function () {

    //assert that a buffer of bytes exist to receive from WASM
    
    //assert that required classes to decode and read the buffer exists
    //assert that various variables related to a state object for DOM manipulation
    // are configured or fallback to some valid configuration

    //cycle instructions from buffer endless loop
        
        //decode OperationCode
        //switch on OperationCode
            //case: instruction to end decoding
                //break cycle
            //case: instruction to perform JS operation
                //perform JS operation on state
                //break switch
            //case: instruction to change state of vuguRender
                //update state
                //break switch
            //default:
                //throw error for invalid OperationCode
                //return
            
}