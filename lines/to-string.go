
//
// Copyright (c) 2017-present Kary Foundation, Inc. All rights reserved.
//   Author: Pouya Kary <k@karyfoundation.org>
//

//
// ─── SETUP ──────────────────────────────────────────────────────────────────────
//

    package lines

    import "strings"

//
// ─── TO STRING ──────────────────────────────────────────────────────────────────
//

    // ToString converts a `Text` type into string
    func ( text Lines ) ToString( ) string {
        return strings.Join( text, "\n" )
    }

// ────────────────────────────────────────────────────────────────────────────────
