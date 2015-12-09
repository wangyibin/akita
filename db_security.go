package akita

type GrantedAuthority interface {
}

type UserDetails interface {
	Authorities() []GrantedAuthority
	Password() string
	Username() string
	IsAccountNotExpired() bool
	IsAccountNonLocked() bool
	IsCredentialsNonExpired() bool
	IsEnabled() bool
}

type UserDetailsService interface {
	LoadUserByUsername(username string) UserDetails
}
