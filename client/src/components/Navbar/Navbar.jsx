import React from 'react'
import './Navbar.css'

const Navbar = () => {
  return (

    <nav className='navbar'>
        <div className='navbar-left'>
           
          
            
            <ul className='nav-links'>
                <li> <a href="/" className='logo'> Lunar Cloud Gaming </a></li>
                <li><a href="#">Home</a></li>
                <li><a href="#">Library</a></li>
                <li><a href="#">Playlist</a></li>
                <li><a href="#">Couch</a></li>
                <li><a href="#">Broadcast</a></li>
                <li><a href="#">Search</a></li>
            </ul>
        </div>
        <div className='navbar-right'>
            <form action="">
                <button>Sign In</button>
            </form>



        </div>

    </nav>
  )
}

export default Navbar