import { IncomingMessage } from 'http'
import getConfig from 'next/config'
import cookie from 'cookie'
import jwt from 'jsonwebtoken'

import { logIn } from './redux'
import { userMapper } from './transform'
import { RoleBasedEntity } from './interfaces'

const { publicRuntimeConfig } = getConfig()

export const validateEmail = (email: string): boolean =>
  email != '' && !!email.match(/.+@.+/)

export const validatePassword = (password: string): boolean =>
  password != '' && password.length >= 6

export const validateDisplayName = (name: string): boolean =>
  /^([\p{Alphabetic}\p{Mark}\p{Decimal_Number}\p{Connector_Punctuation}\p{Join_Control} _-]{2,18})$/u.exec(
    name,
  ) !== null

// const GUEST_ROLE = 0
// const BANNED_ROLE = 1
// const USER_ROLE = 2
const ADMIN_ROLE = 3

export const isAdmin = (entity: RoleBasedEntity | undefined): boolean =>
  entity !== undefined && entity.role >= ADMIN_ROLE

const sessionCookieName =
  publicRuntimeConfig.SESSION_COOKIE_NAME || 'session_token'

export const parseSessionFromContext = (
  request: IncomingMessage | undefined,
  store: any,
) => {
  if (!request) {
    return
  }

  const cookies = cookie.parse(request.headers.cookie || '')
  const sessionCookie = cookies[sessionCookieName]
  const decoded: any = jwt.decode(sessionCookie)

  if (decoded && decoded.User && decoded.exp) {
    store.dispatch(
      logIn({
        expiresAt: decoded.exp,
        user: userMapper.fromRaw(decoded.User),
      }),
    )
  }
}
