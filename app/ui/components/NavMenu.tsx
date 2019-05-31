import React, { useEffect, useState } from 'react'
import Link from 'next/link'
import styled from 'styled-components'
import { connect } from 'react-redux'
import { State } from '../../store'
import { User } from '../../user/interfaces'
import { RankingRegistration } from '../../ranking/interfaces'
import NewLogFormModal from '../../ranking/components/NewLogFormModal'
import Router from 'next/router'

const StyledNav = styled.nav`
  display: flex;
  align-items: center;
`

const NavLink = styled.a`
  padding: 10px;
  display: block;
`

const LoggedInNavigation = ({
  registration,
}: {
  registration: RankingRegistration | undefined
}) => {
  const [open, setOpen] = useState(false)

  return (
    <>
      {registration && (
        <>
          <NavLink href="#" onClick={() => setOpen(true)}>
            Submit pages
          </NavLink>
          <NewLogFormModal
            isOpen={open}
            onCancel={() => setOpen(false)}
            onSuccess={() => {
              setOpen(false)
              if (Router.asPath) {
                Router.push(Router.asPath)
              }
            }}
          />
        </>
      )}
      <Link href="/sign-out">
        <NavLink href="">Sign out</NavLink>
      </Link>
    </>
  )
}

const LoggedOutNavigation = () => (
  <>
    <Link href="/sign-in">
      <NavLink href="">Sign in</NavLink>
    </Link>
    <Link href="/register">
      <NavLink href="">Register</NavLink>
    </Link>
  </>
)

interface Props {
  user: User | undefined
  registration: RankingRegistration | undefined
}

const NavMenu = ({ user, registration }: Props) => {
  const [hasMounted, setHasMounted] = useState(false)

  useEffect(() => {
    setHasMounted(true)
  }, [])

  if (!hasMounted) {
    return null
  }

  return (
    <StyledNav>
      {user ? (
        <LoggedInNavigation registration={registration} />
      ) : (
        <LoggedOutNavigation />
      )}
    </StyledNav>
  )
}

const mapStateToProps = (state: State) => ({
  user: state.session.user,
  registration: state.ranking.registration,
})

export default connect(mapStateToProps)(NavMenu)
