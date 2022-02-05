import React from 'react'
import './Pricing.css'

const Pricing = () => {
    return (
        <React.Fragment>
            <section className="content-container">
         

                
                <ul className="price">
                    <li className="col-header" style={{backgroundColor:'#37387a'}}>Pro</li>
                    <li className="grey">$ 24.99 / year</li>
                    <li>Unlimited Consultation</li>
                    <li>10 - 25 Users</li>
                    <li>7 days Audit logs</li>
                    <li>3 days SLA support</li>
                    <li className="grey"></li>
                </ul>
                

              
           
            </section>
        </React.Fragment>
    )
}

export default Pricing;