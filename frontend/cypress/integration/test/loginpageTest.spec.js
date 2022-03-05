it('Launching', () => {
    cy.visit('http://localhost:3000/login')
    cy.contains('StUni')
      .should('be.visible')
    cy.contains('Home')
      .should('be.visible')
  })