package core

type Server struct {
	/**store all txs that are not included in any blocks
	 * tx must be validated before add into MemPool
	 */
	MemPool map[hash]*Transaction
}
