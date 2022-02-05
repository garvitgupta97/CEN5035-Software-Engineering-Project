import React from 'react'


import Pricing from '../misc/Pricing'
import Header from '../misc/Header'
import MenuBar from '../navigations/MenuBar'


const Home = () => {
    return (
        <React.Fragment>
            <MenuBar/>
            <Header />            
            <Pricing />
         
         
        </React.Fragment>
    )
}

export default Home;
