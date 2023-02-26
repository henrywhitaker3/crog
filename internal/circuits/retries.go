package circuits

import "context"

type Effector func(context.Context) (any, error)

func Retry(e Effector, tries int) Effector {
	return func(ctx context.Context) (any, error) {
		var out any
		var err error
		for i := 0; i < tries; i++ {
			out, err = e(ctx)
			if err == nil {
				return out, err
			}

			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			default:
			}
		}
		return out, err
	}
}
