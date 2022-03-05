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


    cy.contains('StUni')
      .should('be.visible')
    cy.contains('Home')
      .should('be.visible')
  })