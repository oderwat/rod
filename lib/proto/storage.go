// This file is generated by "./lib/proto/generate"

package proto

/*

Storage

*/

// StorageStorageType Enum of possible storage types.
type StorageStorageType string

const (
	// StorageStorageTypeAppcache enum const
	StorageStorageTypeAppcache StorageStorageType = "appcache"

	// StorageStorageTypeCookies enum const
	StorageStorageTypeCookies StorageStorageType = "cookies"

	// StorageStorageTypeFileSystems enum const
	StorageStorageTypeFileSystems StorageStorageType = "file_systems"

	// StorageStorageTypeIndexeddb enum const
	StorageStorageTypeIndexeddb StorageStorageType = "indexeddb"

	// StorageStorageTypeLocalStorage enum const
	StorageStorageTypeLocalStorage StorageStorageType = "local_storage"

	// StorageStorageTypeShaderCache enum const
	StorageStorageTypeShaderCache StorageStorageType = "shader_cache"

	// StorageStorageTypeWebsql enum const
	StorageStorageTypeWebsql StorageStorageType = "websql"

	// StorageStorageTypeServiceWorkers enum const
	StorageStorageTypeServiceWorkers StorageStorageType = "service_workers"

	// StorageStorageTypeCacheStorage enum const
	StorageStorageTypeCacheStorage StorageStorageType = "cache_storage"

	// StorageStorageTypeAll enum const
	StorageStorageTypeAll StorageStorageType = "all"

	// StorageStorageTypeOther enum const
	StorageStorageTypeOther StorageStorageType = "other"
)

// StorageUsageForType Usage for a storage type.
type StorageUsageForType struct {

	// StorageType Name of storage type.
	StorageType StorageStorageType `json:"storageType"`

	// Usage Storage usage (bytes).
	Usage float64 `json:"usage"`
}

// StorageTrustTokens (experimental) Pair of issuer origin and number of available (signed, but not used) Trust
// Tokens from that issuer.
type StorageTrustTokens struct {

	// IssuerOrigin ...
	IssuerOrigin string `json:"issuerOrigin"`

	// Count ...
	Count float64 `json:"count"`
}

// StorageClearDataForOrigin Clears storage for origin.
type StorageClearDataForOrigin struct {

	// Origin Security origin.
	Origin string `json:"origin"`

	// StorageTypes Comma separated list of StorageType to clear.
	StorageTypes string `json:"storageTypes"`
}

// ProtoReq name
func (m StorageClearDataForOrigin) ProtoReq() string { return "Storage.clearDataForOrigin" }

// Call sends the request
func (m StorageClearDataForOrigin) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// StorageGetCookies Returns all browser cookies.
type StorageGetCookies struct {

	// BrowserContextID (optional) Browser context to use when called on the browser endpoint.
	BrowserContextID BrowserBrowserContextID `json:"browserContextId,omitempty"`
}

// ProtoReq name
func (m StorageGetCookies) ProtoReq() string { return "Storage.getCookies" }

// Call the request
func (m StorageGetCookies) Call(c Client) (*StorageGetCookiesResult, error) {
	var res StorageGetCookiesResult
	return &res, call(m.ProtoReq(), m, &res, c)
}

// StorageGetCookiesResult Returns all browser cookies.
type StorageGetCookiesResult struct {

	// Cookies Array of cookie objects.
	Cookies []*NetworkCookie `json:"cookies"`
}

// StorageSetCookies Sets given cookies.
type StorageSetCookies struct {

	// Cookies Cookies to be set.
	Cookies []*NetworkCookieParam `json:"cookies"`

	// BrowserContextID (optional) Browser context to use when called on the browser endpoint.
	BrowserContextID BrowserBrowserContextID `json:"browserContextId,omitempty"`
}

// ProtoReq name
func (m StorageSetCookies) ProtoReq() string { return "Storage.setCookies" }

// Call sends the request
func (m StorageSetCookies) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// StorageClearCookies Clears cookies.
type StorageClearCookies struct {

	// BrowserContextID (optional) Browser context to use when called on the browser endpoint.
	BrowserContextID BrowserBrowserContextID `json:"browserContextId,omitempty"`
}

// ProtoReq name
func (m StorageClearCookies) ProtoReq() string { return "Storage.clearCookies" }

// Call sends the request
func (m StorageClearCookies) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// StorageGetUsageAndQuota Returns usage and quota in bytes.
type StorageGetUsageAndQuota struct {

	// Origin Security origin.
	Origin string `json:"origin"`
}

// ProtoReq name
func (m StorageGetUsageAndQuota) ProtoReq() string { return "Storage.getUsageAndQuota" }

// Call the request
func (m StorageGetUsageAndQuota) Call(c Client) (*StorageGetUsageAndQuotaResult, error) {
	var res StorageGetUsageAndQuotaResult
	return &res, call(m.ProtoReq(), m, &res, c)
}

// StorageGetUsageAndQuotaResult Returns usage and quota in bytes.
type StorageGetUsageAndQuotaResult struct {

	// Usage Storage usage (bytes).
	Usage float64 `json:"usage"`

	// Quota Storage quota (bytes).
	Quota float64 `json:"quota"`

	// OverrideActive Whether or not the origin has an active storage quota override
	OverrideActive bool `json:"overrideActive"`

	// UsageBreakdown Storage usage per type (bytes).
	UsageBreakdown []*StorageUsageForType `json:"usageBreakdown"`
}

// StorageOverrideQuotaForOrigin (experimental) Override quota for the specified origin
type StorageOverrideQuotaForOrigin struct {

	// Origin Security origin.
	Origin string `json:"origin"`

	// QuotaSize (optional) The quota size (in bytes) to override the original quota with.
	// If this is called multiple times, the overridden quota will be equal to
	// the quotaSize provided in the final call. If this is called without
	// specifying a quotaSize, the quota will be reset to the default value for
	// the specified origin. If this is called multiple times with different
	// origins, the override will be maintained for each origin until it is
	// disabled (called without a quotaSize).
	QuotaSize float64 `json:"quotaSize,omitempty"`
}

// ProtoReq name
func (m StorageOverrideQuotaForOrigin) ProtoReq() string { return "Storage.overrideQuotaForOrigin" }

// Call sends the request
func (m StorageOverrideQuotaForOrigin) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// StorageTrackCacheStorageForOrigin Registers origin to be notified when an update occurs to its cache storage list.
type StorageTrackCacheStorageForOrigin struct {

	// Origin Security origin.
	Origin string `json:"origin"`
}

// ProtoReq name
func (m StorageTrackCacheStorageForOrigin) ProtoReq() string {
	return "Storage.trackCacheStorageForOrigin"
}

// Call sends the request
func (m StorageTrackCacheStorageForOrigin) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// StorageTrackIndexedDBForOrigin Registers origin to be notified when an update occurs to its IndexedDB.
type StorageTrackIndexedDBForOrigin struct {

	// Origin Security origin.
	Origin string `json:"origin"`
}

// ProtoReq name
func (m StorageTrackIndexedDBForOrigin) ProtoReq() string { return "Storage.trackIndexedDBForOrigin" }

// Call sends the request
func (m StorageTrackIndexedDBForOrigin) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// StorageUntrackCacheStorageForOrigin Unregisters origin from receiving notifications for cache storage.
type StorageUntrackCacheStorageForOrigin struct {

	// Origin Security origin.
	Origin string `json:"origin"`
}

// ProtoReq name
func (m StorageUntrackCacheStorageForOrigin) ProtoReq() string {
	return "Storage.untrackCacheStorageForOrigin"
}

// Call sends the request
func (m StorageUntrackCacheStorageForOrigin) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// StorageUntrackIndexedDBForOrigin Unregisters origin from receiving notifications for IndexedDB.
type StorageUntrackIndexedDBForOrigin struct {

	// Origin Security origin.
	Origin string `json:"origin"`
}

// ProtoReq name
func (m StorageUntrackIndexedDBForOrigin) ProtoReq() string {
	return "Storage.untrackIndexedDBForOrigin"
}

// Call sends the request
func (m StorageUntrackIndexedDBForOrigin) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// StorageGetTrustTokens (experimental) Returns the number of stored Trust Tokens per issuer for the
// current browsing context.
type StorageGetTrustTokens struct {
}

// ProtoReq name
func (m StorageGetTrustTokens) ProtoReq() string { return "Storage.getTrustTokens" }

// Call the request
func (m StorageGetTrustTokens) Call(c Client) (*StorageGetTrustTokensResult, error) {
	var res StorageGetTrustTokensResult
	return &res, call(m.ProtoReq(), m, &res, c)
}

// StorageGetTrustTokensResult (experimental) Returns the number of stored Trust Tokens per issuer for the
// current browsing context.
type StorageGetTrustTokensResult struct {

	// Tokens ...
	Tokens []*StorageTrustTokens `json:"tokens"`
}

// StorageCacheStorageContentUpdated A cache's contents have been modified.
type StorageCacheStorageContentUpdated struct {

	// Origin Origin to update.
	Origin string `json:"origin"`

	// CacheName Name of cache in origin.
	CacheName string `json:"cacheName"`
}

// ProtoEvent name
func (evt StorageCacheStorageContentUpdated) ProtoEvent() string {
	return "Storage.cacheStorageContentUpdated"
}

// StorageCacheStorageListUpdated A cache has been added/deleted.
type StorageCacheStorageListUpdated struct {

	// Origin Origin to update.
	Origin string `json:"origin"`
}

// ProtoEvent name
func (evt StorageCacheStorageListUpdated) ProtoEvent() string {
	return "Storage.cacheStorageListUpdated"
}

// StorageIndexedDBContentUpdated The origin's IndexedDB object store has been modified.
type StorageIndexedDBContentUpdated struct {

	// Origin Origin to update.
	Origin string `json:"origin"`

	// DatabaseName Database to update.
	DatabaseName string `json:"databaseName"`

	// ObjectStoreName ObjectStore to update.
	ObjectStoreName string `json:"objectStoreName"`
}

// ProtoEvent name
func (evt StorageIndexedDBContentUpdated) ProtoEvent() string {
	return "Storage.indexedDBContentUpdated"
}

// StorageIndexedDBListUpdated The origin's IndexedDB database list has been modified.
type StorageIndexedDBListUpdated struct {

	// Origin Origin to update.
	Origin string `json:"origin"`
}

// ProtoEvent name
func (evt StorageIndexedDBListUpdated) ProtoEvent() string {
	return "Storage.indexedDBListUpdated"
}
