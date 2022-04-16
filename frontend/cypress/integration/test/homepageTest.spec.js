import posts from "../mock-data/getPostsApiMock.js";

it('Launching', () => {
    cy.server()
   
    cy.route({
      method: "GET",
      url: "/posts",
      response: []
    });

    cy.route({
      method: "GET",
      url: "/subreddits",
      response: "[]"
    });

    cy.visit('http://localhost:3000/')

    // Home 
    cy.contains('Home')
      .should('be.visible')
      cy.log('Home button present')

      // Logo
      cy.get('.d-inline-block')
      .should('be.visible')
      cy.log('Logo present')
    
      

      // chakra present and theme change working
      cy.get('.css-a2c8bh > svg > path')
      .should('be.visible') 
      cy.log('chakra present')
      cy.get('.css-a2c8bh > svg > path').click()
      cy.log('dark chakra button working')
      cy.log('theme changing')

      cy.get('.css-1t0t7oo')
      .should('be.visible')
      cy.get('.css-1t0t7oo').click()
      cy.log('light chakra button working')
      cy.log('theme changing')

      // Login present
      cy.contains('Login')
      .should('be.visible')
      cy.log('Login button present')

      // routed to login page
      cy.contains('Login').click()
      cy.go(-1)
      cy.log('routing to Login page')

      // Register present
      cy.contains('Register')
      .should('be.visible')
      cy.log('Register button present')

      // routed to register page
      cy.contains('Register').click()
      cy.go(-1)
      cy.log('routing to Register page')

      // page reloading
      cy.reload(true)
      cy.log('page reloading')
  })