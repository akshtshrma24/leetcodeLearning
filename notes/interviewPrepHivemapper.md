```js
'use client'

import {useDebugValue, createContext, useCallback, useContext, useEffect, useReducer, useState, useDeferredValue, useId, useLayoutEffect, useMemo, useSyncExternalStore } from "react";

export default function Home() {

  const [temp, setTemp] = useState(0)
  
  // Use Context
  const characteristic = {
    tall : "tall",
    short : "short"
  }
  const userContext = createContext() 
  function DisplayUser() {
    return (
      <userContext.Provider value={characteristic.tall}>
        <PTag />
      </userContext.Provider>
    )
  }
  function PTag(){
    const context = useContext(userContext);
    return <p> User Is: {context} </p>
  }

  // useEffect
  useEffect(() =>{
    setTemp(20);
  }, [])


  // Use Reducer
  function reducer(state, action){
    switch(action.type){
      case 'increment':
        return {count : state.count + 1}
      case 'decrement':
        return {count : state.count - 1}
      default: 
        throw new Error()
    }
  }

  const [state, dispatch] = useReducer(reducer, {count : 0})

  // useDebugValue
  const isOnline = true;
  useDebugValue(isOnline ? 'Online' : 'Offline');
  // this is viewable in react dev tools


  //useDefferedValue
  const [query, setQuery] = useState('')
  const deffered = useDeferredValue(query)

  // we dont want to recalculate the list variable unless deferred has been updated
  // so we have deferred in []
  const list = useMemo(() =>{
    console.log(query)
    const dataList = []
    for(let i = 0; i < 10000; i++){
      dataList.push(deffered)
    }
    return dataList
  }, [deffered])

   // useExternalSyncStore   
  function subscribe(callBack) {
    window.addEventListener('online', callBack)
    window.addEventListener('offline', callBack)
    return () =>{
      window.addEventListener('online', callBack)
      window.addEventListener('offline', callBack)
    }
  }

  function snapshot() {
    return navigator.onLine
  }

  const online = useSyncExternalStore(subscribe, snapshot)

  // useId
  const passwordID = useId();
  const usernameID = useId();
  const formID = useId();


  // useCallBack 
  // only re-renders when component unmounts, 
  // temp in the brackets so it can not get re-cached until temp is updated
  // until unmounts
  // that said, it can only count to 1, because it is cached at callBackstate as 0
  const [callBackState, setCallBackState] = useState(0)
  const callbackFunction = useCallback(() =>{
    // expensive operation
    // call to api

    setCallBackState(callBackState  + 1)
  }, [temp]) 

  return (
    <>
      <div className="text-white grid-cols-1 gap-6">
        {/* useState useEffect */}
        <div> 
          <button onClick={() => setTemp(temp + 1)}> increment Temp</button>
          <p>useState: {temp} </p>
        </div>

        {/* useContext */}
        <div>
          <DisplayUser />
        </div>

        {/* useID */}
        <div>
          <label htmlFor={formID}>useID: </label>
          <form id={formID}> 
            <input type="text" id={usernameID} placeholder="username"/>
            <input type="password" id={passwordID} placeholder="password"/>
          </form>
        </div>

        {/* useReducer */}
        <div>
          <button onClick={() => dispatch({type : 'increment'})}> + </button>
          <button onClick={() => dispatch({type : 'decrement'})}> - </button>
          {state.count}
        </div>

        {/* useDefferedValue, useMemo */}
        <div>
          <input type="text" className="text-black" value={query} onChange={(e) => setQuery(e.target.value)} />
          <div className="flex">
            {list.map((item) => {
              return <p> {item} </p>
            })}
          </div>
        </div>

        {/* useCallback */}
        <div>
          <button onClick={callbackFunction}> CallBack function</button>
          {callBackState}
        </div>

        {/* Use Sync External Store */}
        <div>
          <p> useSyncExternalStore : {isOnline ? '✅ Online' : '❌ Disconnected'} </p>
        </div>
      </div>
    </>
  );
}
```

Seperated into divs, and comments,
hooks not used and what they do

seInsertionEffect | used to inject dynamic styles
useLayoutEffect | get layout measurements
useRef | similar to useState, but different in that it  doesnt re-render the component Similar to Pass by reference  It can be used to keep track of variables that dont get  reendered 