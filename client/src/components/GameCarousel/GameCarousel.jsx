import React from 'react'
import cyberPunk from '../../assets/CyberPunk/cyberpunk-2077.webp'
import "./GameCarousel.css"


const GameCarousel = () => {
  
  return (
    <section className='mainSection'>
        <h1 className='Title'>Some title</h1>
        <div className='cardCarousel'>
            <picture className='gamePicture'>
                <img src={ cyberPunk } alt="game card" />
            </picture>
            <picture className='gamePicture'>
                <img src={ cyberPunk } alt="game card" />
            </picture>
            <picture className='gamePicture'>
                <img src={ cyberPunk } alt="game card" />
            </picture>
            <picture className='gamePicture'>
                <img src={ cyberPunk } alt="game card" />
            </picture>
            
        </div>
    </section>
  )
}

export default GameCarousel