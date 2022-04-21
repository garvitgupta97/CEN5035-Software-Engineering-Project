it('Launching', () => {
    
    cy.server()

    cy.visit('http://localhost:3000/register')

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
      url: "users",
      response: '{"user":{"id":422,"username":"testusesssr@gmail.com","created_at":"2022-03-05T03:32:17.876Z","updated_at":"2022-03-05T03:32:17.881Z"},"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NDIyLCJpYXQiOjE2NDY0NTExMzd9.LlW8qO6a_ZHvJBZusY3otXlT9ByMst0ssz2EcBQ0UWM"}'
    });

      // Logo
      cy.get('.d-inline-block')
      .should('be.visible')
      cy.log('Logo present')
 
    // Home 
    cy.contains('Home')
      .should('be.visible')
      cy.log('Home button present')


    cy.get('input[id="username-input"]')
    .clear()
    .type('testuser@gmail.com');
    
    cy.get('input[id="password-input"]')
    .clear()
    .type('testuser');

    cy.get('input[id="confirm-password-input"]')
    .clear()
    .type('testuser');

    // Register working
    cy.findByRole('button', { name: /Register/i }).click();
    cy.log('Register present and working')

    cy.contains('500')
    .should('be.visible')
    cy.log('500 for existing user')

    // creating new user
    // cy.get('input[id="username-input"]')
    // .clear()
    // .type('testuser3@gmail.com');
    
    // cy.get('input[id="password-input"]')
    // .clear()
    // .type('testuser3');

    // cy.get('input[id="confirm-password-input"]')
    // .clear()
    // .type('testuser3');
   
 
  })
