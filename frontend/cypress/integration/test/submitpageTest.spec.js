it('Launching', () => {
    cy.server()

    cy.visit('http://localhost:3000/submit')

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


     cy.get('input[id="username-input"]')
    .clear()
    .type('testuser@gmail.com');
    
    cy.get('input[id="password-input"]')
    .clear()
    .type('testuser');
    cy.log('Logging in')

    cy.findByRole('button', { name: /Login/i }).click();

    // cy.get('#menu-button-10').click();

    // cy.contains('testuser@gmail.com', 'Logout').click();
    cy.findByRole('button', { name: /Submit/i }).click();



    cy.contains('Submit')
    .should('be.visible')

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


        cy.get('#field-7')
        .clear()
        .type('Test Post');
        
        cy.get('#field-8')
        .clear()
        .type('Test Post Content');
        cy.log('Post created')

        cy.findByRole('button', { name: /Submit/i }).click();
 

  })
