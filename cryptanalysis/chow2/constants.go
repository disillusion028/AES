package chow2

import (
	"github.com/OpenWhiteBox/primitives/matrix"
)

// mixColumns is each 8x8 block of the binary MixColumns matrix.
var mixColumns = [4][4]matrix.Matrix{
	[4]matrix.Matrix{
		matrix.Matrix{
			matrix.Row{0x80},
			matrix.Row{0x81},
			matrix.Row{0x02},
			matrix.Row{0x84},
			matrix.Row{0x88},
			matrix.Row{0x10},
			matrix.Row{0x20},
			matrix.Row{0x40},
		},
		matrix.Matrix{
			matrix.Row{0x81},
			matrix.Row{0x83},
			matrix.Row{0x06},
			matrix.Row{0x8c},
			matrix.Row{0x98},
			matrix.Row{0x30},
			matrix.Row{0x60},
			matrix.Row{0xc0},
		},
		matrix.Matrix{
			matrix.Row{0x01},
			matrix.Row{0x02},
			matrix.Row{0x04},
			matrix.Row{0x08},
			matrix.Row{0x10},
			matrix.Row{0x20},
			matrix.Row{0x40},
			matrix.Row{0x80},
		},
		matrix.Matrix{
			matrix.Row{0x01},
			matrix.Row{0x02},
			matrix.Row{0x04},
			matrix.Row{0x08},
			matrix.Row{0x10},
			matrix.Row{0x20},
			matrix.Row{0x40},
			matrix.Row{0x80},
		},
	},
	[4]matrix.Matrix{
		matrix.Matrix{
			matrix.Row{0x01},
			matrix.Row{0x02},
			matrix.Row{0x04},
			matrix.Row{0x08},
			matrix.Row{0x10},
			matrix.Row{0x20},
			matrix.Row{0x40},
			matrix.Row{0x80},
		},
		matrix.Matrix{
			matrix.Row{0x80},
			matrix.Row{0x81},
			matrix.Row{0x02},
			matrix.Row{0x84},
			matrix.Row{0x88},
			matrix.Row{0x10},
			matrix.Row{0x20},
			matrix.Row{0x40},
		},
		matrix.Matrix{
			matrix.Row{0x81},
			matrix.Row{0x83},
			matrix.Row{0x06},
			matrix.Row{0x8c},
			matrix.Row{0x98},
			matrix.Row{0x30},
			matrix.Row{0x60},
			matrix.Row{0xc0},
		},
		matrix.Matrix{
			matrix.Row{0x01},
			matrix.Row{0x02},
			matrix.Row{0x04},
			matrix.Row{0x08},
			matrix.Row{0x10},
			matrix.Row{0x20},
			matrix.Row{0x40},
			matrix.Row{0x80},
		},
	},
	[4]matrix.Matrix{
		matrix.Matrix{
			matrix.Row{0x01},
			matrix.Row{0x02},
			matrix.Row{0x04},
			matrix.Row{0x08},
			matrix.Row{0x10},
			matrix.Row{0x20},
			matrix.Row{0x40},
			matrix.Row{0x80},
		},
		matrix.Matrix{
			matrix.Row{0x01},
			matrix.Row{0x02},
			matrix.Row{0x04},
			matrix.Row{0x08},
			matrix.Row{0x10},
			matrix.Row{0x20},
			matrix.Row{0x40},
			matrix.Row{0x80},
		},
		matrix.Matrix{
			matrix.Row{0x80},
			matrix.Row{0x81},
			matrix.Row{0x02},
			matrix.Row{0x84},
			matrix.Row{0x88},
			matrix.Row{0x10},
			matrix.Row{0x20},
			matrix.Row{0x40},
		},
		matrix.Matrix{
			matrix.Row{0x81},
			matrix.Row{0x83},
			matrix.Row{0x06},
			matrix.Row{0x8c},
			matrix.Row{0x98},
			matrix.Row{0x30},
			matrix.Row{0x60},
			matrix.Row{0xc0},
		},
	},
	[4]matrix.Matrix{
		matrix.Matrix{
			matrix.Row{0x81},
			matrix.Row{0x83},
			matrix.Row{0x06},
			matrix.Row{0x8c},
			matrix.Row{0x98},
			matrix.Row{0x30},
			matrix.Row{0x60},
			matrix.Row{0xc0},
		},
		matrix.Matrix{
			matrix.Row{0x01},
			matrix.Row{0x02},
			matrix.Row{0x04},
			matrix.Row{0x08},
			matrix.Row{0x10},
			matrix.Row{0x20},
			matrix.Row{0x40},
			matrix.Row{0x80},
		},
		matrix.Matrix{
			matrix.Row{0x01},
			matrix.Row{0x02},
			matrix.Row{0x04},
			matrix.Row{0x08},
			matrix.Row{0x10},
			matrix.Row{0x20},
			matrix.Row{0x40},
			matrix.Row{0x80},
		},
		matrix.Matrix{
			matrix.Row{0x80},
			matrix.Row{0x81},
			matrix.Row{0x02},
			matrix.Row{0x84},
			matrix.Row{0x88},
			matrix.Row{0x10},
			matrix.Row{0x20},
			matrix.Row{0x40},
		},
	},
}

// unMixColumns is the same as mixColumns, except each block is inverted.
var unMixColumns = [4][4]matrix.Matrix{
	[4]matrix.Matrix{
		matrix.Matrix{
			matrix.Row{0x03},
			matrix.Row{0x04},
			matrix.Row{0x09},
			matrix.Row{0x11},
			matrix.Row{0x20},
			matrix.Row{0x40},
			matrix.Row{0x80},
			matrix.Row{0x01},
		},
		matrix.Matrix{
			matrix.Row{0xfe},
			matrix.Row{0x03},
			matrix.Row{0x07},
			matrix.Row{0xf0},
			matrix.Row{0x1f},
			matrix.Row{0x3f},
			matrix.Row{0x7f},
			matrix.Row{0xff},
		},
		matrix.Matrix{
			matrix.Row{0x01},
			matrix.Row{0x02},
			matrix.Row{0x04},
			matrix.Row{0x08},
			matrix.Row{0x10},
			matrix.Row{0x20},
			matrix.Row{0x40},
			matrix.Row{0x80},
		},
		matrix.Matrix{
			matrix.Row{0x01},
			matrix.Row{0x02},
			matrix.Row{0x04},
			matrix.Row{0x08},
			matrix.Row{0x10},
			matrix.Row{0x20},
			matrix.Row{0x40},
			matrix.Row{0x80},
		},
	},
	[4]matrix.Matrix{
		matrix.Matrix{
			matrix.Row{0x01},
			matrix.Row{0x02},
			matrix.Row{0x04},
			matrix.Row{0x08},
			matrix.Row{0x10},
			matrix.Row{0x20},
			matrix.Row{0x40},
			matrix.Row{0x80},
		},
		matrix.Matrix{
			matrix.Row{0x03},
			matrix.Row{0x04},
			matrix.Row{0x09},
			matrix.Row{0x11},
			matrix.Row{0x20},
			matrix.Row{0x40},
			matrix.Row{0x80},
			matrix.Row{0x01},
		},
		matrix.Matrix{
			matrix.Row{0xfe},
			matrix.Row{0x03},
			matrix.Row{0x07},
			matrix.Row{0xf0},
			matrix.Row{0x1f},
			matrix.Row{0x3f},
			matrix.Row{0x7f},
			matrix.Row{0xff},
		},
		matrix.Matrix{
			matrix.Row{0x01},
			matrix.Row{0x02},
			matrix.Row{0x04},
			matrix.Row{0x08},
			matrix.Row{0x10},
			matrix.Row{0x20},
			matrix.Row{0x40},
			matrix.Row{0x80},
		},
	},
	[4]matrix.Matrix{
		matrix.Matrix{
			matrix.Row{0x01},
			matrix.Row{0x02},
			matrix.Row{0x04},
			matrix.Row{0x08},
			matrix.Row{0x10},
			matrix.Row{0x20},
			matrix.Row{0x40},
			matrix.Row{0x80},
		},
		matrix.Matrix{
			matrix.Row{0x01},
			matrix.Row{0x02},
			matrix.Row{0x04},
			matrix.Row{0x08},
			matrix.Row{0x10},
			matrix.Row{0x20},
			matrix.Row{0x40},
			matrix.Row{0x80},
		},
		matrix.Matrix{
			matrix.Row{0x03},
			matrix.Row{0x04},
			matrix.Row{0x09},
			matrix.Row{0x11},
			matrix.Row{0x20},
			matrix.Row{0x40},
			matrix.Row{0x80},
			matrix.Row{0x01},
		},
		matrix.Matrix{
			matrix.Row{0xfe},
			matrix.Row{0x03},
			matrix.Row{0x07},
			matrix.Row{0xf0},
			matrix.Row{0x1f},
			matrix.Row{0x3f},
			matrix.Row{0x7f},
			matrix.Row{0xff},
		},
	},
	[4]matrix.Matrix{
		matrix.Matrix{
			matrix.Row{0xfe},
			matrix.Row{0x03},
			matrix.Row{0x07},
			matrix.Row{0xf0},
			matrix.Row{0x1f},
			matrix.Row{0x3f},
			matrix.Row{0x7f},
			matrix.Row{0xff},
		},
		matrix.Matrix{
			matrix.Row{0x01},
			matrix.Row{0x02},
			matrix.Row{0x04},
			matrix.Row{0x08},
			matrix.Row{0x10},
			matrix.Row{0x20},
			matrix.Row{0x40},
			matrix.Row{0x80},
		},
		matrix.Matrix{
			matrix.Row{0x01},
			matrix.Row{0x02},
			matrix.Row{0x04},
			matrix.Row{0x08},
			matrix.Row{0x10},
			matrix.Row{0x20},
			matrix.Row{0x40},
			matrix.Row{0x80},
		},
		matrix.Matrix{
			matrix.Row{0x03},
			matrix.Row{0x04},
			matrix.Row{0x09},
			matrix.Row{0x11},
			matrix.Row{0x20},
			matrix.Row{0x40},
			matrix.Row{0x80},
			matrix.Row{0x01},
		},
	},
}
