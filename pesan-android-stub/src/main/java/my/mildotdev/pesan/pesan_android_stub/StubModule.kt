package my.mildotdev.pesan.pesan_android_stub

import android.content.Context
import android.net.Uri
import android.os.Build
import android.util.Log
import androidx.annotation.RequiresApi
import androidx.credentials.Credential
import androidx.credentials.CredentialManager
import androidx.credentials.GetCredentialRequest
import androidx.credentials.GetPasswordOption
import androidx.credentials.GetPublicKeyCredentialOption
import androidx.credentials.PrepareGetCredentialResponse
import androidx.lifecycle.viewmodel.CreationExtras
import androidx.work.OneTimeWorkRequestBuilder
import com.github.miltsm.pesan_grpc_stubs.OperationHour
import com.github.miltsm.pesan_grpc_stubs.ProtectedGrpcKt
import com.github.miltsm.pesan_grpc_stubs.PublicGrpcKt
import com.github.miltsm.pesan_grpc_stubs.PublicKeyOptions
import com.github.miltsm.pesan_grpc_stubs.RefreshReply
import com.github.miltsm.pesan_grpc_stubs.UserSession
import com.github.miltsm.pesan_grpc_stubs.UuId
import com.github.miltsm.pesan_grpc_stubs.reAuthPasswordRequest
import com.github.miltsm.pesan_grpc_stubs.refreshRequest
import com.github.miltsm.pesan_grpc_stubs.verifyPublicKeyRequest
import com.google.protobuf.ByteString
import com.google.protobuf.Timestamp
import com.google.protobuf.empty
import com.google.protobuf.timestamp
import dagger.Binds
import dagger.Module
import dagger.hilt.InstallIn
import dagger.hilt.android.qualifiers.ApplicationContext
import dagger.hilt.components.SingletonComponent
import io.grpc.CallCredentials
import io.grpc.CallOptions
import io.grpc.Channel
import io.grpc.ClientCall
import io.grpc.ClientInterceptor
import io.grpc.ForwardingClientCall
import io.grpc.ForwardingClientCallListener.SimpleForwardingClientCallListener
import io.grpc.ManagedChannelBuilder
import io.grpc.Metadata
import io.grpc.Metadata.ASCII_STRING_MARSHALLER
import io.grpc.MethodDescriptor
import io.grpc.Status
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.asExecutor
import kotlinx.coroutines.launch
import kotlinx.coroutines.withContext
import kotlinx.datetime.Clock
import kotlinx.datetime.DateTimeUnit
import kotlinx.datetime.Instant
import kotlinx.datetime.LocalTime
import kotlinx.datetime.TimeZone
import kotlinx.datetime.minus
import my.mildotdev.pesan.android_stub.BuildConfig
import my.mildotdev.pesan.pesan_android_stub.SessionRepositoryImpl.Companion.PesanReAuthentication
import java.util.concurrent.Executor
import java.util.concurrent.TimeUnit
import javax.inject.Inject
import javax.inject.Qualifier
import kotlin.uuid.ExperimentalUuidApi
import kotlin.uuid.Uuid

@Qualifier
@Retention(AnnotationRetention.BINARY)
annotation class PesanService

@Qualifier
@Retention(AnnotationRetention.BINARY)
annotation class ProtectedPesanService

@Module
@InstallIn(SingletonComponent::class)
abstract class SessionModule {
    @Binds
    internal abstract fun bindsSessionRepository(impl: SessionRepositoryImpl): SessionRepository
}

// NOTE: android client just use this stub to make calls,
// instead letting them to have control over url, ssl & etc.
object PesanStub {
    fun getStub(): PublicGrpcKt.PublicCoroutineStub {
        val uri =
            Uri.parse("${BuildConfig.BACKEND_SCHEME}://${BuildConfig.BACKEND_HOST}:${BuildConfig.BACKEND_PORT}")
        val channel = let {
//            val creds = TlsChannelCredentials
//                .newBuilder()
//                .requireFakeFeature()
//                .build()
            val bldr = ManagedChannelBuilder.forAddress(uri.host, uri.port)
//            val builder = Grpc.newChannelBuilderForAddress(
//                uri.host,
//                uri.port,
//                InsecureChannelCredentials.create()
//            )
            // TODO: mTLS config; server knows client, client knows server
            if (uri.scheme == "https")
                bldr.useTransportSecurity()
            else
                bldr.usePlaintext()

            bldr.intercept()
            bldr.executor(Dispatchers.IO.asExecutor()).build()
        }

        return PublicGrpcKt.PublicCoroutineStub(channel)
    }

    fun getProtectedStub(context: Context): ProtectedGrpcKt.ProtectedCoroutineStub {
        val uri =
            Uri.parse("${BuildConfig.BACKEND_SCHEME}://${BuildConfig.BACKEND_HOST}:${BuildConfig.BACKEND_PORT}")

        val channel = let {
//            val creds = TlsChannelCredentials
//                .newBuilder()
//                .requireFakeFeature()
//                .build()
            val bldr = ManagedChannelBuilder.forAddress(uri.host, uri.port)
//            val builder = Grpc.newChannelBuilderForAddress(
//                uri.host,
//                uri.port,
//                InsecureChannelCredentials.create()
//            )
            // TODO: mTLS config; server knows client, client knows server
            if (uri.scheme == "https")
                bldr.useTransportSecurity()
            else
                bldr.usePlaintext()

            bldr.executor(Dispatchers.IO.asExecutor()).build()
        }

        return ProtectedGrpcKt.ProtectedCoroutineStub(channel)
    }
}

const val PK_JSON_KEY = "androidx.credentials.BUNDLE_KEY_AUTHENTICATION_RESPONSE_JSON"

class PublicInterceptor : ClientInterceptor {
    companion object {
        const val TAG = "public_stub"
    }

    override fun <ReqT : Any?, RespT : Any?> interceptCall(
        method: MethodDescriptor<ReqT, RespT>?,
        callOptions: CallOptions?,
        next: Channel?
    ): ClientCall<ReqT, RespT> {
        Log.d(TAG, "Sending request for method: ${method?.fullMethodName}")
        return object : ForwardingClientCall.SimpleForwardingClientCall<ReqT, RespT>(
            next?.newCall(
                method,
                callOptions
            )
        ) {
            override fun start(responseListener: Listener<RespT>?, headers: Metadata?) {
                super.start(object : SimpleForwardingClientCallListener<RespT>(responseListener) {
                    override fun onMessage(message: RespT) {
                        Log.d(TAG, "message ->\n$message")
                        super.onMessage(message)
                    }

                    override fun onClose(status: Status?, trailers: Metadata?) {
                        Log.d(TAG, "status ->\n$status")
                        super.onClose(status, trailers)
                    }
                }, headers)
            }
        }
    }
}

class ProtectedInterceptor(val context: Context) : ClientInterceptor {
    companion object {
        const val TAG = "protected_stub"
    }

    override fun <ReqT : Any?, RespT : Any?> interceptCall(
        method: MethodDescriptor<ReqT, RespT>?,
        callOptions: CallOptions?,
        next: Channel?
    ): ClientCall<ReqT, RespT> {
        Log.d(TAG, "Sending request for method: ${method?.fullMethodName}")
        return object : ForwardingClientCall.SimpleForwardingClientCall<ReqT, RespT>(
            next?.newCall(
                method,
                callOptions
            )
        ) {
            override fun start(responseListener: Listener<RespT>?, headers: Metadata?) {
                headers?.put(Metadata.Key.of("Authorization", ASCII_STRING_MARSHALLER), "Bearer ")
                super.start(object : SimpleForwardingClientCallListener<RespT>(responseListener) {
                    override fun onMessage(message: RespT) {
                        Log.d(TAG, "message ->\n$message")
                        super.onMessage(message)
                    }

                    override fun onClose(status: Status?, trailers: Metadata?) {
                        Log.d(TAG, "status ->\n$status")
                        super.onClose(status, trailers)
                    }
                }, headers)
            }
        }
    }
}

// NOTE: for protected calls
class PesanCredential(private val accessToken: String) : CallCredentials() {
    override fun applyRequestMetadata(
        requestInfo: RequestInfo?,
        appExecutor: Executor?,
        applier: MetadataApplier?
    ) {
        val headers = Metadata().apply {
            put(Metadata.Key.of("Authorization", ASCII_STRING_MARSHALLER), "Bearer $accessToken")
        }
        applier?.apply(headers)
    }
}

interface SessionRepository {
    var reauthListener: PesanReAuthentication?
    suspend fun refreshToken(accessToken: String, refreshToken: ByteString): RefreshReply
    suspend fun reAuthViaPublicKey(accessToken: String): PublicKeyOptions
    suspend fun verifyReAuth(accessToken: String, signed: ByteString): RefreshReply
    suspend fun getCredentials(
        userSession: UserSession,
        attestSession: PublicKeyOptions?
    ): Credential

    suspend fun reAuthPassword(accessToken: String, password: ByteString): RefreshReply
    suspend fun getPasswordCredential(): Credential
    suspend fun prepPublicKeyCredential(options: String): PrepareGetCredentialResponse.PendingGetCredentialHandle?
    suspend fun getPublicKeyCredential(pendingGetCredentialHandle: PrepareGetCredentialResponse.PendingGetCredentialHandle): Credential
    suspend fun getPublicKeyCredential(options: String): Credential
}

class SessionRepositoryImpl @Inject constructor(
    @ApplicationContext private val context: Context,
    @ProtectedPesanService private val remoteDataSource: ProtectedGrpcKt.ProtectedCoroutineStub
) : SessionRepository {
    companion object {
        val REPO_KEY = object : CreationExtras.Key<SessionRepository> {}

        interface PesanReAuthentication {
            suspend fun cachePKAuthSession(new: PublicKeyOptions)
        }
    }

    private val dispatcher = Dispatchers.IO
    private val credMgr = CredentialManager.create(context)

    override var reauthListener: PesanReAuthentication? = null

    override suspend fun refreshToken(accessToken: String, refreshToken: ByteString) =
        withContext(dispatcher) {
            remoteDataSource
                .withCallCredentials(PesanCredential(accessToken))
                .refreshSession(refreshRequest {
                    this.token = refreshToken
                }).also {
                    val refreshAt = Instant.fromEpochSeconds(
                        it.accessTokenExpiresAt.seconds,
                        it.accessTokenExpiresAt.nanos
                    ).minus(2, DateTimeUnit.MINUTE, TimeZone.of("Asia/Kuala_Lumpur"))
                    OneTimeWorkRequestBuilder<RefreshWorker>()
                        .setConstraints(sessionConstraints)
                        .setInitialDelay(refreshAt.toEpochMilliseconds(), TimeUnit.MINUTES)
                        .build()
                }
        }

    override suspend fun reAuthViaPublicKey(accessToken: String) =
        withContext(dispatcher) {
            remoteDataSource
                .withCallCredentials(PesanCredential(accessToken))
                .reAuth(empty { })
        }

    override suspend fun verifyReAuth(accessToken: String, signed: ByteString) =
        withContext(dispatcher) {
            remoteDataSource
                .withCallCredentials(PesanCredential(accessToken))
                .verifyPublicKeyReAuth(verifyPublicKeyRequest {
                    this.signed = signed
                })
        }

    override suspend fun getCredentials(
        userSession: UserSession,
        pkOpts: PublicKeyOptions?
    ) =
        withContext(dispatcher) {
            when {
                userSession.totalPasskey == 0 -> GetCredentialRequest(
                    listOf(
                        GetPasswordOption()
                    )
                ).let {
                    credMgr.getCredential(context, it).credential
                }

                (pkOpts == null || Instant.fromEpochSeconds(
                    pkOpts.validUntil.seconds,
                    pkOpts.validUntil.nanos
                ) < Clock.System.now())
                        && Build.VERSION.SDK_INT >= Build.VERSION_CODES.P ->
                    reAuthViaPublicKey(userSession.accessToken.toStringUtf8()).let {
                        launch {
                            reauthListener?.cachePKAuthSession(it)
                        }
                        GetCredentialRequest(
                            listOf(
                                GetPasswordOption(),
                                GetPublicKeyCredentialOption(it.challenge.toStringUtf8())
                            )
                        ).let { cred ->
//                            if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.UPSIDE_DOWN_CAKE)
//                                credMgr.prepareGetCredential(cred).pendingGetCredentialHandle
//                            else
                            credMgr.getCredential(context, cred).credential
                        }
                    }

                pkOpts?.challenge != null && Build.VERSION.SDK_INT >= Build.VERSION_CODES.P -> GetCredentialRequest(
                    listOf(
                        GetPasswordOption(),
                        GetPublicKeyCredentialOption(pkOpts.challenge.toStringUtf8())
                    )
                ).let { cred ->
//                    if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.UPSIDE_DOWN_CAKE)
//                        credMgr.prepareGetCredential(cred).pendingGetCredentialHandle
//                    else
                    credMgr.getCredential(context, cred).credential
                }

                else -> GetCredentialRequest(
                    listOf(
                        GetPasswordOption()
                    )
                ).let {
                    credMgr.getCredential(context, it).credential
                }
            }
        }

    override suspend fun reAuthPassword(accessToken: String, password: ByteString) =
        withContext(dispatcher) {
            remoteDataSource
                .withCallCredentials(PesanCredential(accessToken))
                .reAuthWithPassword(reAuthPasswordRequest {
                    this.password = password
                })
        }

    override suspend fun getPasswordCredential() = credMgr.getCredential(
        context,
        GetCredentialRequest(listOf(GetPasswordOption()))
    ).credential

    @RequiresApi(Build.VERSION_CODES.UPSIDE_DOWN_CAKE)
    override suspend fun prepPublicKeyCredential(options: String): PrepareGetCredentialResponse.PendingGetCredentialHandle? =
        credMgr.prepareGetCredential(
            GetCredentialRequest(
                listOf(
                    GetPasswordOption(),
                    GetPublicKeyCredentialOption(options)
                )
            )
        ).pendingGetCredentialHandle

    @RequiresApi(Build.VERSION_CODES.UPSIDE_DOWN_CAKE)
    override suspend fun getPublicKeyCredential(pendingGetCredentialHandle: PrepareGetCredentialResponse.PendingGetCredentialHandle) =
        credMgr.getCredential(context, pendingGetCredentialHandle).credential

    override suspend fun getPublicKeyCredential(options: String): Credential =
        credMgr.getCredential(
            context,
            GetCredentialRequest(listOf(GetPasswordOption(), GetPublicKeyCredentialOption(options)))
        ).credential
}

@OptIn(ExperimentalUuidApi::class)
fun UuId.uuid() = Uuid.parse(this.id.toStringUtf8())

@OptIn(ExperimentalUuidApi::class)
fun Uuid.pb() = ByteString.copyFromUtf8(this.toString())

fun Instant.pb() = timestamp {
    this.nanos = this@pb.nanosecondsOfSecond
    this.seconds = this@pb.epochSeconds
}

fun Timestamp.instant() = Instant.fromEpochSeconds(this.seconds, this.nanos)

//fun Instant.localTime() = toLocalDateTime(
//    TimeZone.of("Asia/Kuala_Lumpur")
//).let { LocalTime(it.hour, it.minute) }

fun OperationHour.localTime() = LocalTime(hour, minute)

fun String.pb() = ByteString.copyFromUtf8(this)