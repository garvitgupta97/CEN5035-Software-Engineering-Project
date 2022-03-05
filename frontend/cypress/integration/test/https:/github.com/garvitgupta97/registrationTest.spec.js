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


    cy.get('input[id="username-input"]')
    .clear()
    .type('testuser@gmail.com');
    
    cy.get('input[id="password-input"]')
    .clear()
    .type('testuser');
    
  })
