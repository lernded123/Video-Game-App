import gameTitle from "../../../assets/CyberPunk/CPLogo.jpg"
import React from 'react'

const GameTitle = () => {
  return (
    <section>
        <div>
            <picture>
                <img src={gameTitle} alt="" />
            </picture>
        </div>
    </section>

  )
}

export default GameTitle