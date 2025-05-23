// Generated by `wit-bindgen-wrpc-go` 0.11.0. DO NOT EDIT!
package me_gcp_cloudrun_job_admin

import (
	bytes "bytes"
	context "context"
	errors "errors"
	fmt "fmt"
	mattilsynet__me_gcp_cloudrun_job_admin__types "github.com/Mattilsynet/map-me-gcp-cloudrunjob/bindings/mattilsynet/me_gcp_cloudrun_job_admin/types"
	io "io"
	slog "log/slog"
	wrpc "wrpc.io/go"
)

type ManagedEnvironmentGcpManifest = mattilsynet__me_gcp_cloudrun_job_admin__types.ManagedEnvironmentGcpManifest
type Error = mattilsynet__me_gcp_cloudrun_job_admin__types.Error
type Handler interface {
	Update(ctx__ context.Context, manifest *mattilsynet__me_gcp_cloudrun_job_admin__types.ManagedEnvironmentGcpManifest) (*wrpc.Result[ManagedEnvironmentGcpManifest, Error], error)
	Get(ctx__ context.Context, manifest *mattilsynet__me_gcp_cloudrun_job_admin__types.ManagedEnvironmentGcpManifest) (*wrpc.Result[ManagedEnvironmentGcpManifest, Error], error)
	Delete(ctx__ context.Context, manifest *mattilsynet__me_gcp_cloudrun_job_admin__types.ManagedEnvironmentGcpManifest) (*wrpc.Result[ManagedEnvironmentGcpManifest, Error], error)
}

func ServeInterface(s wrpc.Server, h Handler) (stop func() error, err error) {
	stops := make([]func() error, 0, 3)
	stop = func() error {
		for _, stop := range stops {
			if err := stop(); err != nil {
				return err
			}
		}
		return nil
	}

	stop0, err := s.Serve("mattilsynet:me-gcp-cloudrun-job-admin/me-gcp-cloudrun-job-admin@0.1.0", "update", func(ctx context.Context, w wrpc.IndexWriteCloser, r wrpc.IndexReadCloser) {
		defer func() {
			if err := w.Close(); err != nil {
				slog.DebugContext(ctx, "failed to close writer", "instance", "mattilsynet:me-gcp-cloudrun-job-admin/me-gcp-cloudrun-job-admin@0.1.0", "name", "update", "err", err)
			}
		}()
		slog.DebugContext(ctx, "reading parameter", "i", 0)
		p0, err := func() (*ManagedEnvironmentGcpManifest, error) {
			v, err := func(r wrpc.IndexReadCloser, path ...uint32) (*mattilsynet__me_gcp_cloudrun_job_admin__types.ManagedEnvironmentGcpManifest, error) {
				v := &mattilsynet__me_gcp_cloudrun_job_admin__types.ManagedEnvironmentGcpManifest{}
				var err error
				slog.Debug("reading field", "name", "bytes")
				v.Bytes, err = func(r interface {
					io.ByteReader
					io.Reader
				}) ([]byte, error) {
					var x uint32
					var s uint
					for i := 0; i < 5; i++ {
						slog.Debug("reading byte list length", "i", i)
						b, err := r.ReadByte()
						if err != nil {
							if i > 0 && err == io.EOF {
								err = io.ErrUnexpectedEOF
							}
							return nil, fmt.Errorf("failed to read byte list length byte: %w", err)
						}
						if s == 28 && b > 0x0f {
							return nil, errors.New("byte list length overflows a 32-bit integer")
						}
						if b < 0x80 {
							x = x | uint32(b)<<s
							if x == 0 {
								return nil, nil
							}
							buf := make([]byte, x)
							slog.Debug("reading byte list contents", "len", x)
							_, err = io.ReadFull(r, buf)
							if err != nil {
								return nil, fmt.Errorf("failed to read byte list contents: %w", err)
							}
							return buf, nil
						}
						x |= uint32(b&0x7f) << s
						s += 7
					}
					return nil, errors.New("byte length overflows a 32-bit integer")
				}(r)
				if err != nil {
					return nil, fmt.Errorf("failed to read `bytes` field: %w", err)
				}
				return v, nil
			}(r, []uint32{0}...)
			return (*ManagedEnvironmentGcpManifest)(v), err
		}()

		if err != nil {
			slog.WarnContext(ctx, "failed to read parameter", "i", 0, "instance", "mattilsynet:me-gcp-cloudrun-job-admin/me-gcp-cloudrun-job-admin@0.1.0", "name", "update", "err", err)
			if err := r.Close(); err != nil {
				slog.ErrorContext(ctx, "failed to close reader", "instance", "mattilsynet:me-gcp-cloudrun-job-admin/me-gcp-cloudrun-job-admin@0.1.0", "name", "update", "err", err)
			}
			return
		}
		slog.DebugContext(ctx, "calling `mattilsynet:me-gcp-cloudrun-job-admin/me-gcp-cloudrun-job-admin@0.1.0.update` handler")
		r0, err := h.Update(ctx, p0)
		if cErr := r.Close(); cErr != nil {
			slog.ErrorContext(ctx, "failed to close reader", "instance", "mattilsynet:me-gcp-cloudrun-job-admin/me-gcp-cloudrun-job-admin@0.1.0", "name", "update", "err", err)
		}
		if err != nil {
			slog.WarnContext(ctx, "failed to handle invocation", "instance", "mattilsynet:me-gcp-cloudrun-job-admin/me-gcp-cloudrun-job-admin@0.1.0", "name", "update", "err", err)
			return
		}

		var buf bytes.Buffer
		writes := make(map[uint32]func(wrpc.IndexWriter) error, 1)

		write0, err := func(v *wrpc.Result[ManagedEnvironmentGcpManifest, Error], w interface {
			io.ByteWriter
			io.Writer
		}) (func(wrpc.IndexWriter) error, error) {
			switch {
			case v.Ok == nil && v.Err == nil:
				return nil, errors.New("both result variants cannot be nil")
			case v.Ok != nil && v.Err != nil:
				return nil, errors.New("exactly one result variant must non-nil")

			case v.Ok != nil:
				slog.Debug("writing `result::ok` status byte")
				if err := w.WriteByte(0); err != nil {
					return nil, fmt.Errorf("failed to write `result::ok` status byte: %w", err)
				}
				slog.Debug("writing `result::ok` payload")
				write, err := (v.Ok).WriteToIndex(w)
				if err != nil {
					return nil, fmt.Errorf("failed to write `result::ok` payload: %w", err)
				}
				if write != nil {
					return write, nil
				}
				return nil, nil
			default:
				slog.Debug("writing `result::err` status byte")
				if err := w.WriteByte(1); err != nil {
					return nil, fmt.Errorf("failed to write `result::err` status byte: %w", err)
				}
				slog.Debug("writing `result::err` payload")
				write, err := (v.Err).WriteToIndex(w)
				if err != nil {
					return nil, fmt.Errorf("failed to write `result::err` payload: %w", err)
				}
				if write != nil {
					return write, nil
				}
				return nil, nil
			}
		}(r0, &buf)
		if err != nil {
			slog.WarnContext(ctx, "failed to write result value", "i", 0, "instance", "mattilsynet:me-gcp-cloudrun-job-admin/me-gcp-cloudrun-job-admin@0.1.0", "name", "update", "err", err)
			return
		}
		if write0 != nil {
			writes[0] = write0
		}
		slog.DebugContext(ctx, "transmitting `mattilsynet:me-gcp-cloudrun-job-admin/me-gcp-cloudrun-job-admin@0.1.0.update` result")
		_, err = w.Write(buf.Bytes())
		if err != nil {
			slog.WarnContext(ctx, "failed to write result", "instance", "mattilsynet:me-gcp-cloudrun-job-admin/me-gcp-cloudrun-job-admin@0.1.0", "name", "update", "err", err)
			return
		}
		if len(writes) > 0 {
			for index, write := range writes {
				_ = write
				switch index {
				case 0:
					w, err := w.Index(0)
					if err != nil {
						slog.ErrorContext(ctx, "failed to index result writer", "instance", "mattilsynet:me-gcp-cloudrun-job-admin/me-gcp-cloudrun-job-admin@0.1.0", "name", "update", "err", err)
						return
					}
					write := write
					go func() {
						if err := write(w); err != nil {
							slog.WarnContext(ctx, "failed to write nested result value", "instance", "mattilsynet:me-gcp-cloudrun-job-admin/me-gcp-cloudrun-job-admin@0.1.0", "name", "update", "err", err)
						}
					}()
				}
			}
		}
	})
	if err != nil {
		return nil, fmt.Errorf("failed to serve `mattilsynet:me-gcp-cloudrun-job-admin/me-gcp-cloudrun-job-admin@0.1.0.update`: %w", err)
	}
	stops = append(stops, stop0)

	stop1, err := s.Serve("mattilsynet:me-gcp-cloudrun-job-admin/me-gcp-cloudrun-job-admin@0.1.0", "get", func(ctx context.Context, w wrpc.IndexWriteCloser, r wrpc.IndexReadCloser) {
		defer func() {
			if err := w.Close(); err != nil {
				slog.DebugContext(ctx, "failed to close writer", "instance", "mattilsynet:me-gcp-cloudrun-job-admin/me-gcp-cloudrun-job-admin@0.1.0", "name", "get", "err", err)
			}
		}()
		slog.DebugContext(ctx, "reading parameter", "i", 0)
		p0, err := func() (*ManagedEnvironmentGcpManifest, error) {
			v, err := func(r wrpc.IndexReadCloser, path ...uint32) (*mattilsynet__me_gcp_cloudrun_job_admin__types.ManagedEnvironmentGcpManifest, error) {
				v := &mattilsynet__me_gcp_cloudrun_job_admin__types.ManagedEnvironmentGcpManifest{}
				var err error
				slog.Debug("reading field", "name", "bytes")
				v.Bytes, err = func(r interface {
					io.ByteReader
					io.Reader
				}) ([]byte, error) {
					var x uint32
					var s uint
					for i := 0; i < 5; i++ {
						slog.Debug("reading byte list length", "i", i)
						b, err := r.ReadByte()
						if err != nil {
							if i > 0 && err == io.EOF {
								err = io.ErrUnexpectedEOF
							}
							return nil, fmt.Errorf("failed to read byte list length byte: %w", err)
						}
						if s == 28 && b > 0x0f {
							return nil, errors.New("byte list length overflows a 32-bit integer")
						}
						if b < 0x80 {
							x = x | uint32(b)<<s
							if x == 0 {
								return nil, nil
							}
							buf := make([]byte, x)
							slog.Debug("reading byte list contents", "len", x)
							_, err = io.ReadFull(r, buf)
							if err != nil {
								return nil, fmt.Errorf("failed to read byte list contents: %w", err)
							}
							return buf, nil
						}
						x |= uint32(b&0x7f) << s
						s += 7
					}
					return nil, errors.New("byte length overflows a 32-bit integer")
				}(r)
				if err != nil {
					return nil, fmt.Errorf("failed to read `bytes` field: %w", err)
				}
				return v, nil
			}(r, []uint32{0}...)
			return (*ManagedEnvironmentGcpManifest)(v), err
		}()

		if err != nil {
			slog.WarnContext(ctx, "failed to read parameter", "i", 0, "instance", "mattilsynet:me-gcp-cloudrun-job-admin/me-gcp-cloudrun-job-admin@0.1.0", "name", "get", "err", err)
			if err := r.Close(); err != nil {
				slog.ErrorContext(ctx, "failed to close reader", "instance", "mattilsynet:me-gcp-cloudrun-job-admin/me-gcp-cloudrun-job-admin@0.1.0", "name", "get", "err", err)
			}
			return
		}
		slog.DebugContext(ctx, "calling `mattilsynet:me-gcp-cloudrun-job-admin/me-gcp-cloudrun-job-admin@0.1.0.get` handler")
		r0, err := h.Get(ctx, p0)
		if cErr := r.Close(); cErr != nil {
			slog.ErrorContext(ctx, "failed to close reader", "instance", "mattilsynet:me-gcp-cloudrun-job-admin/me-gcp-cloudrun-job-admin@0.1.0", "name", "get", "err", err)
		}
		if err != nil {
			slog.WarnContext(ctx, "failed to handle invocation", "instance", "mattilsynet:me-gcp-cloudrun-job-admin/me-gcp-cloudrun-job-admin@0.1.0", "name", "get", "err", err)
			return
		}

		var buf bytes.Buffer
		writes := make(map[uint32]func(wrpc.IndexWriter) error, 1)

		write0, err := func(v *wrpc.Result[ManagedEnvironmentGcpManifest, Error], w interface {
			io.ByteWriter
			io.Writer
		}) (func(wrpc.IndexWriter) error, error) {
			switch {
			case v.Ok == nil && v.Err == nil:
				return nil, errors.New("both result variants cannot be nil")
			case v.Ok != nil && v.Err != nil:
				return nil, errors.New("exactly one result variant must non-nil")

			case v.Ok != nil:
				slog.Debug("writing `result::ok` status byte")
				if err := w.WriteByte(0); err != nil {
					return nil, fmt.Errorf("failed to write `result::ok` status byte: %w", err)
				}
				slog.Debug("writing `result::ok` payload")
				write, err := (v.Ok).WriteToIndex(w)
				if err != nil {
					return nil, fmt.Errorf("failed to write `result::ok` payload: %w", err)
				}
				if write != nil {
					return write, nil
				}
				return nil, nil
			default:
				slog.Debug("writing `result::err` status byte")
				if err := w.WriteByte(1); err != nil {
					return nil, fmt.Errorf("failed to write `result::err` status byte: %w", err)
				}
				slog.Debug("writing `result::err` payload")
				write, err := (v.Err).WriteToIndex(w)
				if err != nil {
					return nil, fmt.Errorf("failed to write `result::err` payload: %w", err)
				}
				if write != nil {
					return write, nil
				}
				return nil, nil
			}
		}(r0, &buf)
		if err != nil {
			slog.WarnContext(ctx, "failed to write result value", "i", 0, "instance", "mattilsynet:me-gcp-cloudrun-job-admin/me-gcp-cloudrun-job-admin@0.1.0", "name", "get", "err", err)
			return
		}
		if write0 != nil {
			writes[0] = write0
		}
		slog.DebugContext(ctx, "transmitting `mattilsynet:me-gcp-cloudrun-job-admin/me-gcp-cloudrun-job-admin@0.1.0.get` result")
		_, err = w.Write(buf.Bytes())
		if err != nil {
			slog.WarnContext(ctx, "failed to write result", "instance", "mattilsynet:me-gcp-cloudrun-job-admin/me-gcp-cloudrun-job-admin@0.1.0", "name", "get", "err", err)
			return
		}
		if len(writes) > 0 {
			for index, write := range writes {
				_ = write
				switch index {
				case 0:
					w, err := w.Index(0)
					if err != nil {
						slog.ErrorContext(ctx, "failed to index result writer", "instance", "mattilsynet:me-gcp-cloudrun-job-admin/me-gcp-cloudrun-job-admin@0.1.0", "name", "get", "err", err)
						return
					}
					write := write
					go func() {
						if err := write(w); err != nil {
							slog.WarnContext(ctx, "failed to write nested result value", "instance", "mattilsynet:me-gcp-cloudrun-job-admin/me-gcp-cloudrun-job-admin@0.1.0", "name", "get", "err", err)
						}
					}()
				}
			}
		}
	})
	if err != nil {
		return nil, fmt.Errorf("failed to serve `mattilsynet:me-gcp-cloudrun-job-admin/me-gcp-cloudrun-job-admin@0.1.0.get`: %w", err)
	}
	stops = append(stops, stop1)

	stop2, err := s.Serve("mattilsynet:me-gcp-cloudrun-job-admin/me-gcp-cloudrun-job-admin@0.1.0", "delete", func(ctx context.Context, w wrpc.IndexWriteCloser, r wrpc.IndexReadCloser) {
		defer func() {
			if err := w.Close(); err != nil {
				slog.DebugContext(ctx, "failed to close writer", "instance", "mattilsynet:me-gcp-cloudrun-job-admin/me-gcp-cloudrun-job-admin@0.1.0", "name", "delete", "err", err)
			}
		}()
		slog.DebugContext(ctx, "reading parameter", "i", 0)
		p0, err := func() (*ManagedEnvironmentGcpManifest, error) {
			v, err := func(r wrpc.IndexReadCloser, path ...uint32) (*mattilsynet__me_gcp_cloudrun_job_admin__types.ManagedEnvironmentGcpManifest, error) {
				v := &mattilsynet__me_gcp_cloudrun_job_admin__types.ManagedEnvironmentGcpManifest{}
				var err error
				slog.Debug("reading field", "name", "bytes")
				v.Bytes, err = func(r interface {
					io.ByteReader
					io.Reader
				}) ([]byte, error) {
					var x uint32
					var s uint
					for i := 0; i < 5; i++ {
						slog.Debug("reading byte list length", "i", i)
						b, err := r.ReadByte()
						if err != nil {
							if i > 0 && err == io.EOF {
								err = io.ErrUnexpectedEOF
							}
							return nil, fmt.Errorf("failed to read byte list length byte: %w", err)
						}
						if s == 28 && b > 0x0f {
							return nil, errors.New("byte list length overflows a 32-bit integer")
						}
						if b < 0x80 {
							x = x | uint32(b)<<s
							if x == 0 {
								return nil, nil
							}
							buf := make([]byte, x)
							slog.Debug("reading byte list contents", "len", x)
							_, err = io.ReadFull(r, buf)
							if err != nil {
								return nil, fmt.Errorf("failed to read byte list contents: %w", err)
							}
							return buf, nil
						}
						x |= uint32(b&0x7f) << s
						s += 7
					}
					return nil, errors.New("byte length overflows a 32-bit integer")
				}(r)
				if err != nil {
					return nil, fmt.Errorf("failed to read `bytes` field: %w", err)
				}
				return v, nil
			}(r, []uint32{0}...)
			return (*ManagedEnvironmentGcpManifest)(v), err
		}()

		if err != nil {
			slog.WarnContext(ctx, "failed to read parameter", "i", 0, "instance", "mattilsynet:me-gcp-cloudrun-job-admin/me-gcp-cloudrun-job-admin@0.1.0", "name", "delete", "err", err)
			if err := r.Close(); err != nil {
				slog.ErrorContext(ctx, "failed to close reader", "instance", "mattilsynet:me-gcp-cloudrun-job-admin/me-gcp-cloudrun-job-admin@0.1.0", "name", "delete", "err", err)
			}
			return
		}
		slog.DebugContext(ctx, "calling `mattilsynet:me-gcp-cloudrun-job-admin/me-gcp-cloudrun-job-admin@0.1.0.delete` handler")
		r0, err := h.Delete(ctx, p0)
		if cErr := r.Close(); cErr != nil {
			slog.ErrorContext(ctx, "failed to close reader", "instance", "mattilsynet:me-gcp-cloudrun-job-admin/me-gcp-cloudrun-job-admin@0.1.0", "name", "delete", "err", err)
		}
		if err != nil {
			slog.WarnContext(ctx, "failed to handle invocation", "instance", "mattilsynet:me-gcp-cloudrun-job-admin/me-gcp-cloudrun-job-admin@0.1.0", "name", "delete", "err", err)
			return
		}

		var buf bytes.Buffer
		writes := make(map[uint32]func(wrpc.IndexWriter) error, 1)

		write0, err := func(v *wrpc.Result[ManagedEnvironmentGcpManifest, Error], w interface {
			io.ByteWriter
			io.Writer
		}) (func(wrpc.IndexWriter) error, error) {
			switch {
			case v.Ok == nil && v.Err == nil:
				return nil, errors.New("both result variants cannot be nil")
			case v.Ok != nil && v.Err != nil:
				return nil, errors.New("exactly one result variant must non-nil")

			case v.Ok != nil:
				slog.Debug("writing `result::ok` status byte")
				if err := w.WriteByte(0); err != nil {
					return nil, fmt.Errorf("failed to write `result::ok` status byte: %w", err)
				}
				slog.Debug("writing `result::ok` payload")
				write, err := (v.Ok).WriteToIndex(w)
				if err != nil {
					return nil, fmt.Errorf("failed to write `result::ok` payload: %w", err)
				}
				if write != nil {
					return write, nil
				}
				return nil, nil
			default:
				slog.Debug("writing `result::err` status byte")
				if err := w.WriteByte(1); err != nil {
					return nil, fmt.Errorf("failed to write `result::err` status byte: %w", err)
				}
				slog.Debug("writing `result::err` payload")
				write, err := (v.Err).WriteToIndex(w)
				if err != nil {
					return nil, fmt.Errorf("failed to write `result::err` payload: %w", err)
				}
				if write != nil {
					return write, nil
				}
				return nil, nil
			}
		}(r0, &buf)
		if err != nil {
			slog.WarnContext(ctx, "failed to write result value", "i", 0, "instance", "mattilsynet:me-gcp-cloudrun-job-admin/me-gcp-cloudrun-job-admin@0.1.0", "name", "delete", "err", err)
			return
		}
		if write0 != nil {
			writes[0] = write0
		}
		slog.DebugContext(ctx, "transmitting `mattilsynet:me-gcp-cloudrun-job-admin/me-gcp-cloudrun-job-admin@0.1.0.delete` result")
		_, err = w.Write(buf.Bytes())
		if err != nil {
			slog.WarnContext(ctx, "failed to write result", "instance", "mattilsynet:me-gcp-cloudrun-job-admin/me-gcp-cloudrun-job-admin@0.1.0", "name", "delete", "err", err)
			return
		}
		if len(writes) > 0 {
			for index, write := range writes {
				_ = write
				switch index {
				case 0:
					w, err := w.Index(0)
					if err != nil {
						slog.ErrorContext(ctx, "failed to index result writer", "instance", "mattilsynet:me-gcp-cloudrun-job-admin/me-gcp-cloudrun-job-admin@0.1.0", "name", "delete", "err", err)
						return
					}
					write := write
					go func() {
						if err := write(w); err != nil {
							slog.WarnContext(ctx, "failed to write nested result value", "instance", "mattilsynet:me-gcp-cloudrun-job-admin/me-gcp-cloudrun-job-admin@0.1.0", "name", "delete", "err", err)
						}
					}()
				}
			}
		}
	})
	if err != nil {
		return nil, fmt.Errorf("failed to serve `mattilsynet:me-gcp-cloudrun-job-admin/me-gcp-cloudrun-job-admin@0.1.0.delete`: %w", err)
	}
	stops = append(stops, stop2)
	return stop, nil
}
