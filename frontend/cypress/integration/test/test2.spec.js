import 'cypress-react-selector'
import { mount } from '@cypress/react';
import React from 'react'
import HomePage from '/Users/garvitgupta/Documents/Development/CEN5035-Software-Engineering-Project/frontend/src/components/PostList.js'



it('Home Page Rendering ', () => {
    cy.stub(window, 'fetch')
      .withArgs('http://myapi.com/products')
      .resolves({
        json: cy.stub().resolves({
          products: [
            { id: 1, name: 'First item' },
            { id: 2, name: 'Second item' },
          ],
        }),
      })
    // mount the react component
    mount(<HomePage />)
    cy.contains('First item').should('be.visible')
    cy.get('.product').should('have.length', 2)
  
    // use cypress-react-selector 
    cy.waitForReact(1000, '#cypress-root')
    cy.react('ProductsContainer').should('have.class', 'product-container')
    cy.react('AProduct').should('have.length', 2)
    cy.react('AProduct', { name: 'Second item' })
      .should('be.visible')
      .and('have.text', 'Second item')
    cy.getReact('AProduct', { name: 'Second item' })
      .getProps()
      .should('have.property', 'name')
    cy.getReact('AProduct', { name: 'First item' })
      .getProps('name')
      .should('eq', 'First item')
    cy.getReact('AProduct', { name: 'Second item' })
      .getCurrentState()
      .should('not.empty')
  })