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


    cy.get('input[id="username-input"]')
    .clear()
    .type('testuser@gmail.com');
    
    cy.get('input[id="password-input"]')
    .clear()
    .type('testuser');

    cy.findByRole('button', { name: /Login/i }).click();
      cy.contains('.d-inline-block')
      .should('be.visible')
      cy.contains('Submit')
      .should('be.visible')
    cy.contains('Home')
      .should('be.visible')

  })