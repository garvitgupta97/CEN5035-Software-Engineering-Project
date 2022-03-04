describe('Launch Page', () => {
    it('loads successfully', () => {
        cy.visit('http://localhost:3000')
    

    cy.get('.navbar-nav')
            .should('be.visible')
            .within(() => {
                cy.get('h1')
                    .should('contain.text','My Cool MadLibs')
                cy.get('a')
                    .should('be.visible')
                    .should('contain.text', 'Exit Site')
            })
    })
})