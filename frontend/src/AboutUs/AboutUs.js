import React from 'react'
import MenuBar from '../navigations/MenuBar'

const AboutUs = () => {
    return (
        <React.Fragment>
            <MenuBar/>
            <section className="content-container">
                <div className="textArea"> 
                    <h2>About Us</h2>
                    <p>
                        Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec mattis est dui. In id justo id orci ullamcorper sollicitudin. Pellentesque eget dui elementum risus posuere 
                        mollis. Vestibulum nisl arcu, maximus sed mi eget, elementum tempor augue. Vestibulum euismod rhoncus leo et pharetra. Suspendisse pellentesque dapibus tortor. Praesent 
                        varius sem purus, non consectetur sapien facilisis quis. Sed euismod tortor diam, ac imperdiet ligula posuere eu. Pellentesque molestie sit amet felis vitae dapibus. 
                        Morbi iaculis lacinia condimentum. Aenean malesuada tempor vulputate. Proin nunc elit, dictum tempor sagittis eu, pharetra porta ligula. Suspendisse at lacus eu elit rutrum 
                        luctus eu eget erat. Donec massa dui, finibus in maximus id, finibus at urna.
                    </p>   
                </div>
            </section>
        </React.Fragment>
    )
}

export default AboutUs;