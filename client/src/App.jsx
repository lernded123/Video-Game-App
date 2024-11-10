import { useState } from 'react'
import Navbar from "./components/Navbar/Navbar"
import HeroCarousel from './components/MainCarousel/Carousel'
import GameCarousel from './components/GameCarousel/GameCarousel'

function App() {
  

  return (
    <>
      <Navbar />
      <HeroCarousel />
      <GameCarousel />
      <GameCarousel />
      <GameCarousel />
      <GameCarousel />
      <GameCarousel />
      
    </>
  )
}

export default App
