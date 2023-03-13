package circuits

import "context"

type Effector func(context.Context) (any, error)
