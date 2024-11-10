import React from 'react'
import mainCP from '../../assets/CyberPunk/CPMain.jpg'
import logo from '../../assets/CyberPunk/CPLogo.jpg'
import "./Carousel.css"

const Carousel = () => {


  return (
    <main className='grid-container'>
      <div className='numbered-title' id='CarouselID'>
        <img style={{ width:"25rem" }} src= { logo } alt="Game Logo" />
      </div>
      

      <div className='content'>
        
        <button>Learn More</button>
      </div>

    <div className='dot-indicators flex'>
      <button></button>
      <button></button>
      <button></button>
      <button></button>
      <button></button>
      <button></button>
    </div>

    <picture id="mainImg">
      <img  style={{}} src={ mainCP } alt="main Game screen" />
    </picture>





    </main>
  )
}

export default Carousel



    {/* <div class="dot-indicators flex" role="tablist" aria-label="crew member list">
      <button aria-selected="true" role="tab" aria-controls="commander-tab" tabindex="0" data-image="commander-image"><span class="commander-tab"></span></button>
      <button aria-selected="false" role="tab" aria-controls="mission-tab" tabindex="-1" data-image="mission-image"><span class="mission-spec-tab"></span></button>
      <button aria-selected="false" role="tab" aria-controls="pilot-tab" tabindex="-1" data-image="pilot-image"><span class="pilot-tab"></span></button>
      <button aria-selected="false" role="tab" aria-controls="engineer-tab" tabindex="-1" data-image="engineer-image"><span class="engineer-tab"></span></button>
    </div> */}