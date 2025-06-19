(define-library (jakelyz librobot)
  (export get-mouse-position-x
	  get-mouse-position-y
	  drag-mouse-smooth
          move-mouse
	  move-mouse-smooth
	  move-mouse-relative
	  move-mouse-relative-smooth
	  mouse-scroll
	  mouse-click
	  mouse-toggle
	  keyboard-type-string
	  keyboard-key-press
	  keyboard-key-down
	  keyboard-key-up)
  (import (core)
	  (ypsilon c-types)
	  (ypsilon c-ffi))
  (begin
    (define librobot
      (let ((sysname (architecture-feature 'sysname)))
	(cond ((string-contains sysname "darwin")
	       (load-shared-object "librobot.dylib"))
	      ((string-contains sysname "linux")
	       (load-shared-object "librobot.so"))
	      (else
	       (assertion-violation 'load-shared-object "can not load librobot, unexpected operating system")))))
    
    ;; mouse api functions
    (define get-mouse-position-x (c-function long-long GetMousePositionX ()))
    (define get-mouse-position-y (c-function long-long GetMousePositionY ()))
    (define drag-mouse-smooth (c-function void MoveMouseSmooth (long-long long-long double double long-long)))
    (define move-mouse (c-function void MoveMouse (long-long long-long)))
    (define move-mouse-smooth (c-function void MoveMouseSmooth (long-long long-long double double long-long)))
    (define move-mouse-relative (c-function void MoveMouseRelative (long-long long-long)))
    (define move-mouse-relative-smooth (c-function void MoveMouseRelativeSmooth (long-long long-long double double long-long)))
    (define mouse-scroll (c-function void ScrollMouse (long-long long-long)))
    (define mouse-click (c-function void MouseClick (void*)))
    (define mouse-toggle (c-function void MouseToggle (void* void*)))

    ;; keyboard api functions
    (define keyboard-type-string (c-function void TypeString (void*)))
    (define keyboard-key-press (c-function void KeyPress (void*)))
    (define keyboard-key-down (c-function void KeyDown (void*)))
    (define keyboard-key-up (c-function void KeyUp (void*)))
    
    ) ;; end begin
  ) ;; end define-library
