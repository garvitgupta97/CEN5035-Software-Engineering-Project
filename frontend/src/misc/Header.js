import React from 'react'
import { Link } from "react-router-dom"
import './Header.css'

const Header = () => {
    return (
        <React.Fragment>
            <header className="bg-image">
                <div className="bg-container">
                    <h1>Student Universe</h1>
                </div>
            </header>
        </React.Fragment>
    )
}

export default Header;