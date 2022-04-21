import posts from "../mock-data/getPostsApiMock.js";

it('Launching', () => {
    cy.server()

    cy.visit('http://localhost:3000/login')

    cy.route({
      method: "GET",
      url: "/subreddits",
      response: "[]"
    });
    
    cy.route({
      method: "GET",
      url: "/posts",
      response: []
    });

    cy.route({
      method: "POST",
      url: "/api/signin",
      response: '{"user":{"id":419,"username":"testuser@gmail.com","created_at":"2022-03-05T02:16:12.274Z","updated_at":"2022-03-05T02:16:22.021Z"},"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NDE5LCJpYXQiOjE2NDY0NDY1ODJ9.8JmNEBjP0eJk2v16OM8jFp-6D2PO94T_pcWXsOaqRsg"}'
    });


    // Login
    cy.get('input[id="username-input"]')
    .clear()
    .type('testuser@gmail.com');
    
    cy.get('input[id="password-input"]')
    .clear()
    .type('testuser');

    cy.findByRole('button', { name: /Login/i }).click();

    cy.log('Logging in')


    // Logo
    cy.get('.d-inline-block')
      .should('be.visible')
      cy.log('Logo present')

    // Home 
    cy.contains('Home')
      .should('be.visible')
      cy.log('Home button present')

    

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


     // Submit 
    cy.contains('Submit')
    .should('be.visible')
    cy.log('Submit button present')

    cy.contains('Submit').click()
    cy.go(-1)
    cy.log('Submit button routing')

    // User
    cy.get('#menu-button-6').should('be.visible').click()
    cy.log('User button present and working')
 
    // Login
    cy.get('#menu-list-6-menuitem-4').should('be.visible').click()
    cy.log('Logout button present and working')

    // going back after logout to check
    cy.go(-1)
    cy.log('Still logged out')

    
      
   

  })