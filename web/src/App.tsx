import { useState } from 'react'
import logo from './assets/logo.svg'

import './style/App.css'
import Navigation from "./component/Navigation";

function App() {
  const [count, setCount] = useState(0)

  return (
    <div className="App">
      <Navigation />
    </div>
  )
}

export default App
