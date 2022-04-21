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


    cy.contains('Home')
      .should('be.visible')
      cy.get('.d-inline-block')
      .should('be.visible')

      // chakra present and theme change working
      cy.get('.css-a2c8bh > svg > path')
      .should('be.visible') 
      cy.get('.css-a2c8bh > svg > path').click()
      cy.get('.css-1t0t7oo')
      .should('be.visible')
      cy.get('.css-1t0t7oo').click()

      // Login present
      cy.contains('Login')
      .should('be.visible')

      // routed to login page
      cy.contains('Login').click()
      cy.go(-1)

      // Register present
      cy.contains('Register')
      .should('be.visible')

      // routed to register page
      cy.contains('Register').click()
      cy.go(-1)

      // page reloading
      cy.reload(true)
  })